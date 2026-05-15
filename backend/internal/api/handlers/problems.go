package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/executo/backend/internal/db"
	"github.com/executo/backend/internal/models"
)

// ProblemsHandler handles all problem-related HTTP endpoints.
type ProblemsHandler struct {
	db *db.DB
}

// NewProblemsHandler creates a new ProblemsHandler.
func NewProblemsHandler(database *db.DB) *ProblemsHandler {
	return &ProblemsHandler{db: database}
}

// ── GET /problems ─────────────────────────────

// ListProblems returns a paginated list of problems with optional filtering.
func (h *ProblemsHandler) ListProblems(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	difficulty := r.URL.Query().Get("difficulty")
	search := r.URL.Query().Get("search")
	page := parseIntParam(r.URL.Query().Get("page"), 1)
	pageSize := parseIntParam(r.URL.Query().Get("page_size"), 20)

	// Clamp page size
	if pageSize > 100 {
		pageSize = 100
	}
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	// Validate difficulty
	var difficultyParam interface{}
	if difficulty != "" && difficulty != "all" {
		switch models.Difficulty(difficulty) {
		case models.DifficultyEasy, models.DifficultyMedium, models.DifficultyHard:
			difficultyParam = difficulty
		default:
			writeError(w, http.StatusBadRequest, "invalid difficulty: must be easy, medium, or hard")
			return
		}
	}

	// Null-safe search param
	var searchParam interface{}
	if search != "" {
		searchParam = search
	}

	// Count total matching problems
	var total int
	err := h.db.QueryRowContext(r.Context(), `
		SELECT COUNT(*) FROM problems
		WHERE ($1::VARCHAR IS NULL OR difficulty = $1)
		  AND ($2::VARCHAR IS NULL OR title ILIKE '%' || $2 || '%')
	`, difficultyParam, searchParam).Scan(&total)
	if err != nil {
		log.Printf("Error counting problems: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to count problems")
		return
	}

	// Fetch problems
	rows, err := h.db.QueryContext(r.Context(), `
		SELECT
			id, title, slug, difficulty,
			total_submissions,
			CASE
				WHEN total_submissions = 0 THEN 0.0
				ELSE ROUND((accepted_submissions::NUMERIC / total_submissions) * 100, 1)
			END AS acceptance_rate
		FROM problems
		WHERE ($1::VARCHAR IS NULL OR difficulty = $1)
		  AND ($2::VARCHAR IS NULL OR title ILIKE '%' || $2 || '%')
		ORDER BY id ASC
		LIMIT $3 OFFSET $4
	`, difficultyParam, searchParam, pageSize, offset)
	if err != nil {
		log.Printf("Error listing problems: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to list problems")
		return
	}
	defer rows.Close()

	problems := make([]models.ProblemSummary, 0)
	for rows.Next() {
		var p models.ProblemSummary
		if err := rows.Scan(
			&p.ID, &p.Title, &p.Slug, &p.Difficulty,
			&p.TotalSubmissions, &p.AcceptanceRate,
		); err != nil {
			log.Printf("Error scanning problem row: %v", err)
			continue
		}
		problems = append(problems, p)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating problem rows: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to read problems")
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"data":      problems,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// ── GET /problems/:id ─────────────────────────

// GetProblem returns a single problem by ID.
func (h *ProblemsHandler) GetProblem(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromPath(r.URL.Path, "/problems/")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid problem ID")
		return
	}

	var p models.Problem
	err = h.db.QueryRowContext(r.Context(), `
		SELECT
			id, title, slug, description, difficulty,
			examples, constraints, test_cases, function_signature,
			total_submissions,
			CASE
				WHEN total_submissions = 0 THEN 0.0
				ELSE ROUND((accepted_submissions::NUMERIC / total_submissions) * 100, 1)
			END AS acceptance_rate,
			created_at
		FROM problems
		WHERE id = $1
	`, id).Scan(
		&p.ID, &p.Title, &p.Slug, &p.Description, &p.Difficulty,
		&p.Examples, &p.Constraints, &p.TestCases, &p.FunctionSignature,
		&p.TotalSubmissions, &p.AcceptanceRate, &p.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			writeError(w, http.StatusNotFound, "problem not found")
			return
		}
		log.Printf("Error fetching problem %d: %v", id, err)
		writeError(w, http.StatusInternalServerError, "failed to fetch problem")
		return
	}

	// Don't expose test cases to the client (they're hidden)
	p.TestCases = nil

	writeJSON(w, http.StatusOK, p)
}

// ── GET /problems/slug/:slug ──────────────────

// GetProblemBySlug returns a single problem by its URL slug.
func (h *ProblemsHandler) GetProblemBySlug(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/problems/slug/")
	if slug == "" {
		writeError(w, http.StatusBadRequest, "slug is required")
		return
	}

	var p models.Problem
	err := h.db.QueryRowContext(r.Context(), `
		SELECT
			id, title, slug, description, difficulty,
			examples, constraints, test_cases, function_signature,
			total_submissions, created_at
		FROM problems
		WHERE slug = $1
	`, slug).Scan(
		&p.ID, &p.Title, &p.Slug, &p.Description, &p.Difficulty,
		&p.Examples, &p.Constraints, &p.TestCases, &p.FunctionSignature,
		&p.TotalSubmissions, &p.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			writeError(w, http.StatusNotFound, "problem not found")
			return
		}
		log.Printf("Error fetching problem by slug %q: %v", slug, err)
		writeError(w, http.StatusInternalServerError, "failed to fetch problem")
		return
	}

	p.TestCases = nil
	writeJSON(w, http.StatusOK, p)
}

// ── GET /stats ────────────────────────────────

// GetStats returns platform-wide statistics.
func (h *ProblemsHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	var totalProblems, totalSubmissions, totalAccepted int64

	err := h.db.QueryRowContext(r.Context(), `
		SELECT
			COUNT(*) AS total_problems,
			COALESCE(SUM(total_submissions), 0) AS total_submissions,
			COALESCE(SUM(accepted_submissions), 0) AS total_accepted
		FROM problems
	`).Scan(&totalProblems, &totalSubmissions, &totalAccepted)
	if err != nil {
		log.Printf("Error fetching stats: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to fetch stats")
		return
	}

	var acceptanceRate float64
	if totalSubmissions > 0 {
		acceptanceRate = float64(totalAccepted) / float64(totalSubmissions) * 100
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"total_problems":    totalProblems,
		"total_submissions": totalSubmissions,
		"total_accepted":    totalAccepted,
		"acceptance_rate":   acceptanceRate,
	})
}

// ── Helpers ───────────────────────────────────

// writeJSON writes a JSON response with the given status code.
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

// writeError writes a JSON error response.
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

// parseIntParam parses an integer query parameter with a default value.
func parseIntParam(s string, defaultVal int) int {
	if s == "" {
		return defaultVal
	}
	n, err := strconv.Atoi(s)
	if err != nil || n < 1 {
		return defaultVal
	}
	return n
}

// extractIDFromPath extracts a numeric ID from a URL path.
// e.g. "/problems/42" with prefix "/problems/" returns 42.
func extractIDFromPath(path, prefix string) (int64, error) {
	idStr := strings.TrimPrefix(path, prefix)
	// Remove any trailing path segments
	if idx := strings.Index(idStr, "/"); idx != -1 {
		idStr = idStr[:idx]
	}
	return strconv.ParseInt(idStr, 10, 64)
}
