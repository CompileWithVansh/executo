// Package domain defines core business types and errors for Executo.
// This package has ZERO external dependencies — only stdlib.
package domain

import "fmt"

// ── Error Types ───────────────────────────────
// These allow handlers to map domain errors to HTTP status codes cleanly.

// ErrNotFound indicates a resource was not found.
type ErrNotFound struct {
	Resource string
	ID       interface{}
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("%s %v not found", e.Resource, e.ID)
}

// ErrValidation indicates invalid input.
type ErrValidation struct {
	Field   string
	Message string
}

func (e *ErrValidation) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("validation error: %s — %s", e.Field, e.Message)
	}
	return fmt.Sprintf("validation error: %s", e.Message)
}

// ErrConflict indicates a duplicate or conflicting resource.
type ErrConflict struct {
	Resource string
	Message  string
}

func (e *ErrConflict) Error() string {
	return fmt.Sprintf("conflict: %s — %s", e.Resource, e.Message)
}

// ErrInternal indicates an unexpected internal error.
type ErrInternal struct {
	Cause error
}

func (e *ErrInternal) Error() string {
	return fmt.Sprintf("internal error: %v", e.Cause)
}

func (e *ErrInternal) Unwrap() error {
	return e.Cause
}

// ErrRateLimit indicates too many requests.
type ErrRateLimit struct{}

func (e *ErrRateLimit) Error() string {
	return "rate limit exceeded"
}

// ── Helper Constructors ───────────────────────

func NotFound(resource string, id interface{}) error {
	return &ErrNotFound{Resource: resource, ID: id}
}

func ValidationError(field, message string) error {
	return &ErrValidation{Field: field, Message: message}
}

func InternalError(cause error) error {
	return &ErrInternal{Cause: cause}
}
