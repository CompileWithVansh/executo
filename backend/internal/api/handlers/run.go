// Package handlers — Playground run endpoint (direct code execution without problem context)
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/executo/backend/internal/executor"
)

// RunHandler handles the playground /run endpoint for direct code execution.
type RunHandler struct {
	judge0 *executor.Judge0Client
}

// NewRunHandler creates a new RunHandler.
func NewRunHandler(judge0Client *executor.Judge0Client) *RunHandler {
	return &RunHandler{
		judge0: judge0Client,
	}
}

// RunRequest is the payload for POST /run
type RunRequest struct {
	SourceCode string `json:"source_code"` // base64-encoded
	LanguageID int    `json:"language_id"`
	Stdin      string `json:"stdin"`       // base64-encoded (optional)
}

// ── POST /run ─────────────────────────────────
// Submits code to Judge0 for execution (playground mode, no test cases).
func (h *RunHandler) CreateRun(w http.ResponseWriter, r *http.Request) {
	// Limit request body size to 1MB to prevent memory exhaustion
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var req RunRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.SourceCode == "" {
		writeError(w, http.StatusBadRequest, "source_code is required")
		return
	}
	if req.LanguageID == 0 {
		writeError(w, http.StatusBadRequest, "language_id is required")
		return
	}

	// Decode base64 source code for Judge0
	// The frontend sends base64-encoded code, and Judge0 also expects base64,
	// so we pass it through directly.
	token, err := h.judge0.SubmitRaw(req.SourceCode, req.LanguageID, req.Stdin)
	if err != nil {
		log.Printf("Judge0 run submit error: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to submit code for execution")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}

// ── GET /run/:token ───────────────────────────
// Polls Judge0 for the result of a playground run.
func (h *RunHandler) GetRunResult(w http.ResponseWriter, r *http.Request) {
	// Extract token from path: /run/{token}
	token := strings.TrimPrefix(r.URL.Path, "/run/")
	if token == "" {
		writeError(w, http.StatusBadRequest, "token is required")
		return
	}

	result, err := h.judge0.GetResult(token)
	if err != nil {
		log.Printf("Judge0 get result error for token %s: %v", token, err)
		writeError(w, http.StatusInternalServerError, "failed to get execution result")
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"stdout":         result.Stdout,
		"stderr":         result.Stderr,
		"compile_output": result.CompileOutput,
		"message":        result.Message,
		"status": map[string]interface{}{
			"id":          result.Status.ID,
			"description": result.Status.Description,
		},
		"time":   result.Time,
		"memory": result.Memory,
	})
}
