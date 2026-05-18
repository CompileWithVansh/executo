// Package logger provides structured JSON logging for the Executo backend.
// Each log entry includes timestamp, level, message, request_id, and optional fields.
package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Level represents log severity.
type Level string

const (
	LevelDebug Level = "DEBUG"
	LevelInfo  Level = "INFO"
	LevelWarn  Level = "WARN"
	LevelError Level = "ERROR"
)

// Entry is a single structured log entry.
type Entry struct {
	Timestamp string                 `json:"timestamp"`
	Level     Level                  `json:"level"`
	Message   string                 `json:"message"`
	RequestID string                 `json:"request_id,omitempty"`
	UserID    int64                  `json:"user_id,omitempty"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

// Logger is a structured JSON logger.
type Logger struct {
	minLevel Level
}

// New creates a new Logger. In production, set minLevel to LevelInfo.
func New() *Logger {
	level := LevelInfo
	if os.Getenv("GO_ENV") == "development" {
		level = LevelDebug
	}
	return &Logger{minLevel: level}
}

// Info logs an informational message.
func (l *Logger) Info(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.log(ctx, LevelInfo, msg, fields...)
}

// Warn logs a warning message.
func (l *Logger) Warn(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.log(ctx, LevelWarn, msg, fields...)
}

// Error logs an error message.
func (l *Logger) Error(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.log(ctx, LevelError, msg, fields...)
}

// Debug logs a debug message (only in development).
func (l *Logger) Debug(ctx context.Context, msg string, fields ...map[string]interface{}) {
	l.log(ctx, LevelDebug, msg, fields...)
}

func (l *Logger) log(ctx context.Context, level Level, msg string, fields ...map[string]interface{}) {
	if !l.shouldLog(level) {
		return
	}

	entry := Entry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     level,
		Message:   msg,
	}

	// Extract request ID and user ID from context if available
	if ctx != nil {
		if reqID, ok := ctx.Value(requestIDKey).(string); ok {
			entry.RequestID = reqID
		}
		if userID, ok := ctx.Value(userIDKey).(int64); ok {
			entry.UserID = userID
		}
	}

	// Merge extra fields
	if len(fields) > 0 && fields[0] != nil {
		entry.Fields = fields[0]
	}

	// Output as JSON
	data, err := json.Marshal(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, `{"level":"ERROR","message":"failed to marshal log entry: %v"}`+"\n", err)
		return
	}
	fmt.Fprintln(os.Stdout, string(data))
}

func (l *Logger) shouldLog(level Level) bool {
	levels := map[Level]int{
		LevelDebug: 0,
		LevelInfo:  1,
		LevelWarn:  2,
		LevelError: 3,
	}
	return levels[level] >= levels[l.minLevel]
}

// Context keys (matching middleware package)
type ctxKey string

const (
	requestIDKey ctxKey = "request_id"
	userIDKey    ctxKey = "user_id"
)
