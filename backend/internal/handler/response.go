// Package handler provides HTTP request handlers for the Executo API.
// Handlers are thin — they parse the request, call a service, and write the response.
package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/executo/backend/internal/domain"
)

// ── JSON Response Helpers ─────────────────────

// JSON writes a JSON response with the given status code.
func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

// Error writes a JSON error response, mapping domain errors to HTTP status codes.
func Error(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	message := "internal server error"

	var notFound *domain.ErrNotFound
	var validation *domain.ErrValidation
	var conflict *domain.ErrConflict
	var rateLimit *domain.ErrRateLimit

	switch {
	case errors.As(err, &notFound):
		status = http.StatusNotFound
		message = notFound.Error()
	case errors.As(err, &validation):
		status = http.StatusBadRequest
		message = validation.Error()
	case errors.As(err, &conflict):
		status = http.StatusConflict
		message = conflict.Error()
	case errors.As(err, &rateLimit):
		status = http.StatusTooManyRequests
		message = "too many requests"
	default:
		log.Printf("Internal error: %v", err)
	}

	JSON(w, status, map[string]string{"error": message})
}

// ErrorMsg writes a simple error message with a status code.
func ErrorMsg(w http.ResponseWriter, status int, message string) {
	JSON(w, status, map[string]string{"error": message})
}

// ── Request Parsing Helpers ───────────────────

// PathID extracts a numeric ID from a URL path.
// e.g. "/problems/42" with prefix "/problems/" returns 42.
func PathID(path, prefix string) (int64, error) {
	idStr := strings.TrimPrefix(path, prefix)
	if idx := strings.Index(idStr, "/"); idx != -1 {
		idStr = idStr[:idx]
	}
	return strconv.ParseInt(idStr, 10, 64)
}

// QueryInt returns an integer query parameter with a default value.
func QueryInt(r *http.Request, key string, defaultVal int) int {
	s := r.URL.Query().Get(key)
	if s == "" {
		return defaultVal
	}
	n, err := strconv.Atoi(s)
	if err != nil || n < 1 {
		return defaultVal
	}
	return n
}

// QueryString returns a string query parameter.
func QueryString(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
