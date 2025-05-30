// Package config provides centralized, validated configuration for the Executo backend.
// All environment variables are read once at startup and exposed as a typed struct.
package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all application configuration.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Judge0   Judge0Config
	Auth     AuthConfig
}

// ServerConfig holds HTTP server settings.
type ServerConfig struct {
	Port string
	Env  string // "development" or "production"
}

// DatabaseConfig holds PostgreSQL connection settings.
type DatabaseConfig struct {
	URL             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int // minutes
}

// RedisConfig holds Redis connection settings.
type RedisConfig struct {
	URL string
}

// Judge0Config holds Judge0 API settings.
type Judge0Config struct {
	URL    string
	APIKey string
}

// AuthConfig holds authentication settings.
type AuthConfig struct {
	JWTSecret string
}

// Load reads all configuration from environment variables.
// It returns an error if required variables are missing.
func Load() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Port: getEnv("BACKEND_PORT", "8080"),
			Env:  getEnv("GO_ENV", "development"),
		},
		Database: DatabaseConfig{
			URL:             buildDatabaseURL(),
			MaxOpenConns:    getEnvInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvInt("DB_MAX_IDLE_CONNS", 10),
			ConnMaxLifetime: getEnvInt("DB_CONN_MAX_LIFETIME_MIN", 5),
		},
		Redis: RedisConfig{
			URL: getEnv("REDIS_URL", "redis://localhost:6379"),
		},
		Judge0: Judge0Config{
			URL:    getEnv("JUDGE0_URL", "http://localhost:2358"),
			APIKey: getEnv("JUDGE0_API_KEY", ""),
		},
		Auth: AuthConfig{
			JWTSecret: getEnv("JWT_SECRET", ""),
		},
	}

	// Validate required fields
	if cfg.Database.URL == "" {
		return nil, fmt.Errorf("database URL is required (set POSTGRES_URL or individual POSTGRES_* vars)")
	}
	if cfg.Auth.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	return cfg, nil
}

// buildDatabaseURL constructs a PostgreSQL DSN from environment variables.
func buildDatabaseURL() string {
	if url := os.Getenv("POSTGRES_URL"); url != "" {
		return url
	}
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		getEnv("POSTGRES_USER", "executo"),
		getEnv("POSTGRES_PASSWORD", "executo"),
		getEnv("POSTGRES_HOST", "localhost"),
		getEnv("POSTGRES_PORT", "5432"),
		getEnv("POSTGRES_DB", "executo_db"),
	)
}

// getEnv returns the value of an environment variable or a default.
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

// getEnvInt returns an integer environment variable or a default.
func getEnvInt(key string, defaultVal int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return n
}
