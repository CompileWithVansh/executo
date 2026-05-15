// Package handlers — test case suggestion + admin approval endpoints
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/executo/backend/internal/db"
	"github.com/executo/backend/internal/models"
)

// TestCasesHandler handles suggested test case routes.
type TestCasesHandler struct {
	db *db.DB
}

func NewTestCasesHandler(database *db.DB) *TestCasesHandler {
	return &TestCasesHandler{db: database}
}

// ── POST /problems/:slug/suggest-testcase ─────────────────────────────────
// Anyone can suggest a test case. Saved as status = 'pending'.
func (h *TestCasesHandler) SuggestTestCase(w http.ResponseWriter, r *http.Request) {
	// Extract slug from path: /problems/{slug}/suggest-testcase
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	// parts = ["problems", "{slug}", "suggest-testcase"]
	if len(parts) < 3 {
		jsonError(w, "invalid path", http.StatusBadRequest)
		return
	}
	slug := parts[1]

	// Decode request body
	var req models.SuggestTestCaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Look up problem by slug
	var problemID int64
	err := h.db.QueryRow(
		`SELECT id FROM problems WHERE slug = $1`, slug,
	).Scan(&problemID)
	if err != nil {
		jsonError(w, "problem not found", http.StatusNotFound)
		return
	}

	// Get submitter IP (anonymous attribution)
	ip := r.RemoteAddr
	if fwd := r.Header.Get("X-Forwarded-For"); fwd != "" {
		ip = strings.Split(fwd, ",")[0]
	}

	// Insert suggestion
	var id int64
	err = h.db.QueryRow(
		`INSERT INTO suggested_testcases
		    (problem_id, input, expected_output, note, submitted_by)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id`,
		problemID, req.Input, req.ExpectedOutput, req.Note, ip,
	).Scan(&id)
	if err != nil {
		jsonError(w, "failed to save suggestion", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      id,
		"message": "Test case suggestion submitted. It will be reviewed by an admin.",
	})
}

// ── GET /problems/:slug/testcases ─────────────────────────────────────────
// Returns only approved test cases for a problem (public).
func (h *TestCasesHandler) GetApprovedTestCases(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 {
		jsonError(w, "invalid path", http.StatusBadRequest)
		return
	}
	slug := parts[1]

	rows, err := h.db.Query(
		`SELECT stc.id, stc.problem_id, stc.input, stc.expected_output, stc.note, stc.created_at
		 FROM suggested_testcases stc
		 JOIN problems p ON p.id = stc.problem_id
		 WHERE p.slug = $1 AND stc.status = 'approved'
		 ORDER BY stc.created_at ASC`,
		slug,
	)
	if err != nil {
		jsonError(w, "database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type PublicTestCase struct {
		ID             int64  `json:"id"`
		ProblemID      int64  `json:"problem_id"`
		Input          string `json:"input"`
		ExpectedOutput string `json:"expected_output"`
		Note           string `json:"note,omitempty"`
	}

	results := []PublicTestCase{}
	for rows.Next() {
		var tc PublicTestCase
		var note *string
		var createdAt interface{}
		if err := rows.Scan(&tc.ID, &tc.ProblemID, &tc.Input, &tc.ExpectedOutput, &note, &createdAt); err != nil {
			continue
		}
		if note != nil {
			tc.Note = *note
		}
		results = append(results, tc)
	}

	jsonOK(w, results)
}

// ── GET /admin/testcases ──────────────────────────────────────────────────
// Returns all pending test case suggestions for admin review.
// Optional ?status=pending|approved|rejected filter.
func (h *TestCasesHandler) AdminListTestCases(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if status == "" {
		status = "pending"
	}

	rows, err := h.db.Query(
		`SELECT stc.id, stc.problem_id, p.title, p.slug,
		        stc.input, stc.expected_output, stc.note,
		        stc.status, stc.submitted_by, stc.admin_note,
		        stc.created_at, stc.updated_at
		 FROM suggested_testcases stc
		 JOIN problems p ON p.id = stc.problem_id
		 WHERE stc.status = $1
		 ORDER BY stc.created_at ASC`,
		status,
	)
	if err != nil {
		jsonError(w, "database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type AdminTestCase struct {
		ID             int64  `json:"id"`
		ProblemID      int64  `json:"problem_id"`
		ProblemTitle   string `json:"problem_title"`
		ProblemSlug    string `json:"problem_slug"`
		Input          string `json:"input"`
		ExpectedOutput string `json:"expected_output"`
		Note           string `json:"note,omitempty"`
		Status         string `json:"status"`
		SubmittedBy    string `json:"submitted_by,omitempty"`
		AdminNote      string `json:"admin_note,omitempty"`
		CreatedAt      string `json:"created_at"`
		UpdatedAt      string `json:"updated_at"`
	}

	results := []AdminTestCase{}
	for rows.Next() {
		var tc AdminTestCase
		var note, submittedBy, adminNote *string
		if err := rows.Scan(
			&tc.ID, &tc.ProblemID, &tc.ProblemTitle, &tc.ProblemSlug,
			&tc.Input, &tc.ExpectedOutput, &note,
			&tc.Status, &submittedBy, &adminNote,
			&tc.CreatedAt, &tc.UpdatedAt,
		); err != nil {
			continue
		}
		if note != nil {
			tc.Note = *note
		}
		if submittedBy != nil {
			tc.SubmittedBy = *submittedBy
		}
		if adminNote != nil {
			tc.AdminNote = *adminNote
		}
		results = append(results, tc)
	}

	jsonOK(w, results)
}

// ── PATCH /admin/testcases/:id ────────────────────────────────────────────
// Approve or reject a suggested test case.
func (h *TestCasesHandler) AdminPatchTestCase(w http.ResponseWriter, r *http.Request) {
	// Extract ID from path: /admin/testcases/{id}
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 3 {
		jsonError(w, "invalid path", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		jsonError(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req models.PatchTestCaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		jsonError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.db.Exec(
		`UPDATE suggested_testcases
		 SET status = $1, admin_note = $2, updated_at = NOW()
		 WHERE id = $3`,
		req.Status, req.AdminNote, id,
	)
	if err != nil {
		jsonError(w, "database error", http.StatusInternalServerError)
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		jsonError(w, "test case not found", http.StatusNotFound)
		return
	}

	jsonOK(w, map[string]interface{}{
		"id":     id,
		"status": req.Status,
	})
}

// ── helpers ───────────────────────────────────────────────────────────────

func jsonOK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func jsonError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
