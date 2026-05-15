// Package models defines the core domain types for Executo.
package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// Difficulty represents the difficulty level of a problem.
type Difficulty string

const (
	DifficultyEasy   Difficulty = "easy"
	DifficultyMedium Difficulty = "medium"
	DifficultyHard   Difficulty = "hard"
)

// Example is a single input/output example shown in the problem description.
type Example struct {
	Input       string `json:"input"`
	Output      string `json:"output"`
	Explanation string `json:"explanation,omitempty"`
}

// TestCase is a hidden test case used to judge a submission.
type TestCase struct {
	Input          string `json:"input"`
	ExpectedOutput string `json:"expected_output"`
}

// FunctionSignatures holds the starter code for each supported language.
type FunctionSignatures map[string]string

// JSONArray is a helper type for scanning JSON arrays from PostgreSQL.
type JSONArray[T any] []T

// Value implements the driver.Valuer interface for database writes.
func (j JSONArray[T]) Value() (driver.Value, error) {
	if j == nil {
		return "[]", nil
	}
	b, err := json.Marshal(j)
	if err != nil {
		return nil, fmt.Errorf("marshaling JSONArray: %w", err)
	}
	return string(b), nil
}

// Scan implements the sql.Scanner interface for database reads.
func (j *JSONArray[T]) Scan(src any) error {
	var data []byte
	switch v := src.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	case nil:
		*j = JSONArray[T]{}
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", src)
	}
	return json.Unmarshal(data, j)
}

// JSONMap is a helper type for scanning JSON objects from PostgreSQL.
type JSONMap map[string]string

// Value implements the driver.Valuer interface.
func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return "{}", nil
	}
	b, err := json.Marshal(j)
	if err != nil {
		return nil, fmt.Errorf("marshaling JSONMap: %w", err)
	}
	return string(b), nil
}

// Scan implements the sql.Scanner interface.
func (j *JSONMap) Scan(src any) error {
	var data []byte
	switch v := src.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	case nil:
		*j = JSONMap{}
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", src)
	}
	return json.Unmarshal(data, j)
}

// Problem represents a coding problem in the system.
type Problem struct {
	ID                int64                  `json:"id"`
	Title             string                 `json:"title"`
	Slug              string                 `json:"slug"`
	Description       string                 `json:"description"`
	Difficulty        Difficulty             `json:"difficulty"`
	LCNumber          *int                   `json:"lc_number,omitempty"`
	LCUrl             *string                `json:"lc_url,omitempty"`
	Examples          JSONArray[Example]     `json:"examples"`
	Constraints       JSONArray[string]      `json:"constraints"`
	TestCases         JSONArray[TestCase]    `json:"test_cases"`
	FunctionSignature JSONMap                `json:"function_signature"`
	AcceptanceRate    float64                `json:"acceptance_rate"`
	TotalSubmissions  int64                  `json:"total_submissions"`
	CreatedAt         time.Time              `json:"created_at"`
}

// ProblemSummary is a lightweight version of Problem for list views.
type ProblemSummary struct {
	ID               int64      `json:"id"`
	Title            string     `json:"title"`
	Slug             string     `json:"slug"`
	Difficulty       Difficulty `json:"difficulty"`
	LCNumber         *int       `json:"lc_number,omitempty"`
	LCUrl            *string    `json:"lc_url,omitempty"`
	AcceptanceRate   float64    `json:"acceptance_rate"`
	TotalSubmissions int64      `json:"total_submissions"`
}

// SuggestedTestCase is a user-submitted test case pending admin review.
type SuggestedTestCase struct {
	ID             int64     `json:"id"`
	ProblemID      int64     `json:"problem_id"`
	Input          string    `json:"input"`
	ExpectedOutput string    `json:"expected_output"`
	Note           string    `json:"note,omitempty"`
	Status         string    `json:"status"` // pending | approved | rejected
	SubmittedBy    string    `json:"submitted_by,omitempty"`
	AdminNote      string    `json:"admin_note,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// SuggestTestCaseRequest is the payload for POST /problems/:slug/suggest-testcase
type SuggestTestCaseRequest struct {
	Input          string `json:"input"`
	ExpectedOutput string `json:"expected_output"`
	Note           string `json:"note,omitempty"`
}

func (r *SuggestTestCaseRequest) Validate() error {
	if r.Input == "" {
		return fmt.Errorf("input is required")
	}
	if r.ExpectedOutput == "" {
		return fmt.Errorf("expected_output is required")
	}
	return nil
}

// PatchTestCaseRequest is the payload for PATCH /admin/testcases/:id
type PatchTestCaseRequest struct {
	Status    string `json:"status"`    // approved | rejected
	AdminNote string `json:"admin_note,omitempty"`
}

func (r *PatchTestCaseRequest) Validate() error {
	if r.Status != "approved" && r.Status != "rejected" {
		return fmt.Errorf("status must be 'approved' or 'rejected'")
	}
	return nil
}

// CreateProblemRequest is the payload for creating a new problem.
type CreateProblemRequest struct {
	Title             string             `json:"title"`
	Slug              string             `json:"slug"`
	Description       string             `json:"description"`
	Difficulty        Difficulty         `json:"difficulty"`
	Examples          []Example          `json:"examples"`
	Constraints       []string           `json:"constraints"`
	TestCases         []TestCase         `json:"test_cases"`
	FunctionSignature FunctionSignatures `json:"function_signature"`
}

// Validate checks that required fields are present.
func (r *CreateProblemRequest) Validate() error {
	if r.Title == "" {
		return fmt.Errorf("title is required")
	}
	if r.Slug == "" {
		return fmt.Errorf("slug is required")
	}
	if r.Description == "" {
		return fmt.Errorf("description is required")
	}
	switch r.Difficulty {
	case DifficultyEasy, DifficultyMedium, DifficultyHard:
		// valid
	default:
		return fmt.Errorf("difficulty must be easy, medium, or hard")
	}
	if len(r.TestCases) == 0 {
		return fmt.Errorf("at least one test case is required")
	}
	return nil
}
