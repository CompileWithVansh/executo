 // Package db provides PostgreSQL database connectivity and query helpers.
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB wraps the standard sql.DB with application-specific helpers.
type DB struct {
	*sql.DB
}

// New opens a PostgreSQL connection using the POSTGRES_URL environment variable.
// It configures the connection pool and verifies connectivity with a ping.
func New() (*DB, error) {
	dsn := os.Getenv("POSTGRES_URL")
	if dsn == "" {
		// Build DSN from individual env vars as fallback
		dsn = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			getEnv("POSTGRES_USER", "executo"),
			getEnv("POSTGRES_PASSWORD", "executo"),
			getEnv("POSTGRES_HOST", "localhost"),
			getEnv("POSTGRES_PORT", "5432"),
			getEnv("POSTGRES_DB", "executo_db"),
		)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	// Connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(2 * time.Minute)

	// Verify connectivity
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	log.Println("✓ Connected to PostgreSQL")
	return &DB{db}, nil
}

// MustNew is like New but panics on error. Use in main() only.
func MustNew() *DB {
	db, err := New()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

// WithRetry attempts to connect to the database with exponential backoff.
// Useful when the database container may not be ready immediately.
func WithRetry(maxAttempts int) (*DB, error) {
	var (
		db  *DB
		err error
	)

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = New()
		if err == nil {
			return db, nil
		}

		if attempt < maxAttempts {
			wait := time.Duration(attempt*attempt) * time.Second
			log.Printf("Database connection attempt %d/%d failed: %v. Retrying in %v...",
				attempt, maxAttempts, err, wait)
			time.Sleep(wait)
		}
	}

	return nil, fmt.Errorf("failed to connect after %d attempts: %w", maxAttempts, err)
}

// getEnv returns the value of an environment variable or a default.
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
