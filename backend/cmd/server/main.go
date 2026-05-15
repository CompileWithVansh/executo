// Executo Backend — Main Entry Point
//
// This binary starts:
//   1. The HTTP API server (handles REST requests)
//   2. The Asynq worker (processes submission execution jobs)
//
// Both run concurrently in the same process.
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/executo/backend/internal/api"
	"github.com/executo/backend/internal/db"
	"github.com/executo/backend/internal/executor"
	"github.com/executo/backend/internal/queue"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting Executo backend...")

	// ── Database ──────────────────────────────
	// Retry connection up to 10 times (for Docker Compose startup ordering)
	database, err := db.WithRetry(10)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// ── Queue Client ──────────────────────────
	// Used by HTTP handlers to enqueue submission jobs
	queueClient := queue.NewClient()
	defer queueClient.Close()

	// ── Judge0 Client ─────────────────────────
	judge0Client := executor.NewJudge0Client()
	log.Printf("Judge0 URL: %s", os.Getenv("JUDGE0_URL"))

	// ── Asynq Worker ──────────────────────────
	// Runs in a separate goroutine, processes jobs from Redis
	worker := queue.NewWorker(database, judge0Client)
	go func() {
		if err := worker.Start(); err != nil {
			log.Printf("Worker stopped: %v", err)
		}
	}()

	// ── HTTP Server ───────────────────────────
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}

	router := api.NewRouter(database, queueClient)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine so we can handle shutdown signals
	go func() {
		log.Printf("✓ HTTP server listening on :%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// ── Graceful Shutdown ─────────────────────
	// Wait for SIGINT or SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gracefully...")

	// Give in-flight requests 30 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("HTTP server forced shutdown: %v", err)
	}

	worker.Shutdown()
	log.Println("Server stopped.")
}
