// Package api sets up the HTTP router and all routes for the Executo backend.
package api

import (
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/executo/backend/internal/api/handlers"
	"github.com/executo/backend/internal/api/middleware"
	"github.com/executo/backend/internal/db"
	"github.com/executo/backend/internal/executor"
	"github.com/executo/backend/internal/queue"
)

// NewRouter creates and configures the HTTP router with all routes.
// It wires together handlers, middleware, and dependencies.
func NewRouter(database *db.DB, queueClient *queue.Client) http.Handler {
	// Initialize handlers
	problemsHandler := handlers.NewProblemsHandler(database)
	submissionsHandler := handlers.NewSubmissionsHandler(database, queueClient)
	testCasesHandler := handlers.NewTestCasesHandler(database)
	runHandler := handlers.NewRunHandler(executor.NewJudge0Client())
	authHandler := handlers.NewAuthHandler(database)

	// Create a new ServeMux
	mux := http.NewServeMux()

	// ── Health ──────────────────────────────────
	mux.HandleFunc("/health", handlers.HealthHandler)

	// ── Auth ────────────────────────────────────
	// POST /auth/register — create a new account
	// POST /auth/login    — get JWT token
	// GET  /auth/me       — get current user (requires auth)
	// PUT  /auth/profile  — update bio/avatar (requires auth)
	mux.HandleFunc("/auth/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		authHandler.Register(w, r)
	})

	mux.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		authHandler.Login(w, r)
	})

	mux.HandleFunc("/auth/me", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		middleware.RequireAuth(http.HandlerFunc(authHandler.GetMe)).ServeHTTP(w, r)
	})

	mux.HandleFunc("/auth/profile", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		middleware.RequireAuth(http.HandlerFunc(authHandler.UpdateProfile)).ServeHTTP(w, r)
	})

	// ── Playground Run ──────────────────────────
	// POST /run       — submit code for execution (playground)
	// GET  /run/:token — poll for execution result
	mux.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		runHandler.CreateRun(w, r)
	})

	mux.HandleFunc("/run/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		runHandler.GetRunResult(w, r)
	})

	// ── Problems ────────────────────────────────
	// GET /problems          — list all problems (paginated, filterable)
	// GET /problems/:id      — get problem by ID
	// GET /problems/slug/:slug — get problem by slug
	mux.HandleFunc("/problems", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			problemsHandler.ListProblems(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/problems/slug/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		problemsHandler.GetProblemBySlug(w, r)
	})

	// GET  /problems/:slug/testcases       — approved test cases (public)
	// POST /problems/:slug/suggest-testcase — submit a suggestion
	mux.HandleFunc("/problems/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		switch {
		case strings.HasSuffix(path, "/suggest-testcase"):
			if r.Method != http.MethodPost {
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
				return
			}
			testCasesHandler.SuggestTestCase(w, r)

		case strings.HasSuffix(path, "/testcases"):
			if r.Method != http.MethodGet {
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
				return
			}
			testCasesHandler.GetApprovedTestCases(w, r)

		default:
			if r.Method != http.MethodGet {
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
				return
			}
			problemsHandler.GetProblem(w, r)
		}
	})

	// ── Submissions ──────────────────────────────
	// POST /submissions      — create a new submission
	// GET  /submissions      — list submissions (filter by ?problem_id=)
	// GET  /submissions/:id  — get submission by ID (for polling)
	mux.HandleFunc("/submissions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// Apply stricter rate limiting for submissions
			middleware.SubmissionRateLimit(
				http.HandlerFunc(submissionsHandler.CreateSubmission),
			).ServeHTTP(w, r)
		case http.MethodGet:
			submissionsHandler.ListSubmissions(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/submissions/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		submissionsHandler.GetSubmission(w, r)
	})

	// ── Admin ────────────────────────────────────
	// GET   /admin/testcases      — list pending suggestions (admin only)
	// PATCH /admin/testcases/:id  — approve or reject (admin only)
	mux.HandleFunc("/admin/testcases", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// Require auth + admin role
		middleware.RequireAuth(
			middleware.RequireAdmin(
				http.HandlerFunc(testCasesHandler.AdminListTestCases),
			),
		).ServeHTTP(w, r)
	})

	mux.HandleFunc("/admin/testcases/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// Require auth + admin role
		middleware.RequireAuth(
			middleware.RequireAdmin(
				http.HandlerFunc(testCasesHandler.AdminPatchTestCase),
			),
		).ServeHTTP(w, r)
	})

	// ── Stats & Leaderboard ──────────────────────
	mux.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		problemsHandler.GetStats(w, r)
	})

	mux.HandleFunc("/leaderboard", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		submissionsHandler.GetLeaderboard(w, r)
	})

	// ── Prometheus Metrics ───────────────────────
	// Scraped by Prometheus at /metrics
	mux.Handle("/metrics", promhttp.Handler())

	// ── Apply Global Middleware ──────────────────
	// Order: RequestID → CORS → Rate Limit → Metrics → Router
	globalRateLimiter := middleware.NewRateLimiter(100, 200) // 100 req/s, burst 200

	handler := middleware.RequestID(
		middleware.CORS(
			globalRateLimiter.Limit(
				middleware.Metrics(mux),
			),
		),
	)

	return handler
}
