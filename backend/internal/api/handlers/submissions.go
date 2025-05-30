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
	"github.com/executo/backend/internal/queue"
)

// SubmissionsHandler handles all submission-related HTTP endpoints.
type SubmissionsHandler struct {
	db     *db.DB
	queue  *queue.Client
}

// NewSubmissionsHandler creates a new SubmissionsHandler.
func NewSubmissionsHandler(database *db.DB, queueClient *queue.Client) *SubmissionsHandler {
	return &SubmissionsHandler{
		db:    database,
		queue: queueClient,
	}
}

// ── POST /submissions ─────────────────────────

// CreateSubmission accepts a new code submission, saves it to the DB,
// and enqueues it for execution.
func (h *SubmissionsHandler) CreateSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Limit request body size to 1MB to prevent memory exhaustion
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	// Parse request body
	var req models.CreateSubmissionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	// Validate
	if err := req.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Verify problem exists and get test case count
	var testCasesTotal int
	var testCasesJSON string
	err := h.db.QueryRowContext(r.Context(),
		"SELECT test_cases FROM problems WHERE id = $1",
		req.ProblemID,
	).Scan(&testCasesJSON)
	if err != nil {
		if err == sql.ErrNoRows {
			writeError(w, http.StatusNotFound, "problem not found")
			return
		}
		log.Printf("Error fetching problem %d: %v", req.ProblemID, err)
		writeError(w, http.StatusInternalServerError, "failed to verify problem")
		return
	}

	// Count test cases from JSON
	var testCases []models.TestCase
	if err := json.Unmarshal([]byte(testCasesJSON), &testCases); err == nil {
		testCasesTotal = len(testCases)
	}

	// Create submission record
	var submissionID int64
	err = h.db.QueryRowContext(r.Context(), `
		INSERT INTO submissions (problem_id, language, source_code, status, test_cases_total)
		VALUES ($1, $2, $3, 'pending', $4)
		RETURNING id
	`, req.ProblemID, string(req.Language), req.SourceCode, testCasesTotal).Scan(&submissionID)
	if err != nil {
		log.Printf("Error creating submission: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to create submission")
		return
	}

	// Enqueue for async execution
	if err := h.queue.EnqueueSubmission(submissionID); err != nil {
		log.Printf("Error enqueuing submission #%d: %v", submissionID, err)
		// Don't fail the request — the submission is saved, just not queued
		// A background job could retry this
	}

	writeJSON(w, http.StatusCreated, models.CreateSubmissionResponse{
		ID:      submissionID,
		Status:  models.StatusPending,
		Message: "Submission received and queued for execution",
	})
}

// ── GET /submissions/:id ──────────────────────

// GetSubmission returns a single submission by ID.
// Used by the frontend to poll for results.
func (h *SubmissionsHandler) GetSubmission(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromPath(r.URL.Path, "/submissions/")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid submission ID")
		return
	}

	var s models.Submission
	var verdict, stdout, stderr, compileOutput sql.NullString
	var runtimeMS, memoryKB sql.NullInt64

	err = h.db.QueryRowContext(r.Context(), `
		SELECT
			id, problem_id, language, source_code,
			status, verdict, stdout, stderr, compile_output,
			runtime_ms, memory_kb,
			test_cases_passed, test_cases_total,
			created_at, updated_at
		FROM submissions
		WHERE id = $1
	`, id).Scan(
		&s.ID, &s.ProblemID, &s.Language, &s.SourceCode,
		&s.Status, &verdict, &stdout, &stderr, &compileOutput,
		&runtimeMS, &memoryKB,
		&s.TestCasesPassed, &s.TestCasesTotal,
		&s.CreatedAt, &s.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			writeError(w, http.StatusNotFound, "submission not found")
			return
		}
		log.Printf("Error fetching submission %d: %v", id, err)
		writeError(w, http.StatusInternalServerError, "failed to fetch submission")
		return
	}

	// Map nullable fields
	if verdict.Valid {
		s.Verdict = verdict.String
	}
	if stdout.Valid {
		s.Stdout = stdout.String
	}
	if stderr.Valid {
		s.Stderr = stderr.String
	}
	if compileOutput.Valid {
		s.CompileOutput = compileOutput.String
	}
	if runtimeMS.Valid {
		s.RuntimeMS = &runtimeMS.Int64
	}
	if memoryKB.Valid {
		s.MemoryKB = &memoryKB.Int64
	}

	// Don't expose source code in list views (only in detail)
	writeJSON(w, http.StatusOK, s)
}

// ── GET /submissions?problem_id=:id ──────────

// ListSubmissions returns submissions, optionally filtered by problem_id.
func (h *SubmissionsHandler) ListSubmissions(w http.ResponseWriter, r *http.Request) {
	problemIDStr := r.URL.Query().Get("problem_id")

	var rows *sql.Rows
	var err error

	if problemIDStr != "" {
		problemID, parseErr := parseInt64(problemIDStr)
		if parseErr != nil {
			writeError(w, http.StatusBadRequest, "invalid problem_id")
			return
		}
		rows, err = h.db.QueryContext(r.Context(), `
			SELECT
				id, problem_id, language, status,
				runtime_ms, memory_kb,
				test_cases_passed, test_cases_total,
				created_at
			FROM submissions
			WHERE problem_id = $1
			ORDER BY created_at DESC
			LIMIT 50
		`, problemID)
	} else {
		rows, err = h.db.QueryContext(r.Context(), `
			SELECT
				id, problem_id, language, status,
				runtime_ms, memory_kb,
				test_cases_passed, test_cases_total,
				created_at
			FROM submissions
			ORDER BY created_at DESC
			LIMIT 50
		`)
	}

	if err != nil {
		log.Printf("Error listing submissions: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to list submissions")
		return
	}
	defer rows.Close()

	submissions := make([]map[string]interface{}, 0)
	for rows.Next() {
		var (
			id, problemID                    int64
			language, status                 string
			runtimeMS, memoryKB              sql.NullInt64
			testCasesPassed, testCasesTotal  int
			createdAt                        string
		)

		if err := rows.Scan(
			&id, &problemID, &language, &status,
			&runtimeMS, &memoryKB,
			&testCasesPassed, &testCasesTotal,
			&createdAt,
		); err != nil {
			log.Printf("Error scanning submission row: %v", err)
			continue
		}

		entry := map[string]interface{}{
			"id":                 id,
			"problem_id":         problemID,
			"language":           language,
			"status":             status,
			"test_cases_passed":  testCasesPassed,
			"test_cases_total":   testCasesTotal,
			"created_at":         createdAt,
		}
		if runtimeMS.Valid {
			entry["runtime_ms"] = runtimeMS.Int64
		}
		if memoryKB.Valid {
			entry["memory_kb"] = memoryKB.Int64
		}

		submissions = append(submissions, entry)
	}

	writeJSON(w, http.StatusOK, submissions)
}

// ── GET /leaderboard ─────────────────────────

// GetLeaderboard returns a simple leaderboard based on accepted submissions.
func (h *SubmissionsHandler) GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.QueryContext(r.Context(), `
		SELECT
			p.id,
			p.title,
			p.difficulty,
			COUNT(CASE WHEN s.status = 'accepted' THEN 1 END) AS accepted_count,
			COUNT(*) AS total_count
		FROM problems p
		LEFT JOIN submissions s ON s.problem_id = p.id
		GROUP BY p.id, p.title, p.difficulty
		ORDER BY accepted_count DESC, p.id ASC
		LIMIT 20
	`)
	if err != nil {
		log.Printf("Error fetching leaderboard: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to fetch leaderboard")
		return
	}
	defer rows.Close()

	type entry struct {
		Rank          int     `json:"rank"`
		UserID        string  `json:"user_id"`
		Username      string  `json:"username"`
		ProblemsSolved int    `json:"problems_solved"`
		TotalSubmissions int  `json:"total_submissions"`
		AcceptanceRate float64 `json:"acceptance_rate"`
		Score         int     `json:"score"`
	}

	// For now, return problem-based stats (user auth is out of scope)
	var entries []map[string]interface{}
	rank := 1
	for rows.Next() {
		var (
			id, acceptedCount, totalCount int64
			title, difficulty             string
		)
		if err := rows.Scan(&id, &title, &difficulty, &acceptedCount, &totalCount); err != nil {
			continue
		}

		var rate float64
		if totalCount > 0 {
			rate = float64(acceptedCount) / float64(totalCount) * 100
		}

		entries = append(entries, map[string]interface{}{
			"rank":              rank,
			"user_id":           strings.ToLower(title[:1]),
			"username":          title,
			"problems_solved":   acceptedCount,
			"total_submissions": totalCount,
			"acceptance_rate":   rate,
			"score":             acceptedCount * 100,
		})
		rank++
	}

	if entries == nil {
		entries = []map[string]interface{}{}
	}

	writeJSON(w, http.StatusOK, entries)
}

// ── Helpers ───────────────────────────────────

func parseInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
