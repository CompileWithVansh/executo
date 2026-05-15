// Package handlers contains HTTP request handlers for the Executo API.
package handlers

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

// startTime records when the server started (for uptime calculation).
var startTime = time.Now()

// HealthResponse is the response body for the health check endpoint.
type HealthResponse struct {
	Status    string            `json:"status"`
	Version   string            `json:"version"`
	Uptime    string            `json:"uptime"`
	Timestamp string            `json:"timestamp"`
	Checks    map[string]string `json:"checks"`
	GoVersion string            `json:"go_version"`
}

// HealthHandler handles GET /health
// Returns the service health status and basic diagnostics.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime).Round(time.Second)

	resp := HealthResponse{
		Status:    "ok",
		Version:   "1.0.0",
		Uptime:    uptime.String(),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		GoVersion: runtime.Version(),
		Checks: map[string]string{
			"api": "ok",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
