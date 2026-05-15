package models

import (
	"fmt"
	"time"
)

// Language represents a supported programming language.
type Language string

const (
	LanguagePython3    Language = "python3"
	LanguageJava       Language = "java"
	LanguageCPP        Language = "cpp"
	LanguageJavaScript Language = "javascript"
)

// Judge0LanguageID maps our language identifiers to Judge0 language IDs.
var Judge0LanguageID = map[Language]int{
	LanguagePython3:    71,
	LanguageJava:       62,
	LanguageCPP:        54,
	LanguageJavaScript: 63,
}

// SubmissionStatus represents the current state of a submission.
type SubmissionStatus string

const (
	StatusPending             SubmissionStatus = "pending"
	StatusProcessing          SubmissionStatus = "processing"
	StatusAccepted            SubmissionStatus = "accepted"
	StatusWrongAnswer         SubmissionStatus = "wrong_answer"
	StatusTimeLimitExceeded   SubmissionStatus = "time_limit_exceeded"
	StatusMemoryLimitExceeded SubmissionStatus = "memory_limit_exceeded"
	StatusRuntimeError        SubmissionStatus = "runtime_error"
	StatusCompilationError    SubmissionStatus = "compilation_error"
	StatusInternalError       SubmissionStatus = "internal_error"
)

// Submission represents a code submission by a user.
type Submission struct {
	ID               int64            `json:"id"`
	ProblemID        int64            `json:"problem_id"`
	Language         Language         `json:"language"`
	SourceCode       string           `json:"source_code"`
	Status           SubmissionStatus `json:"status"`
	Verdict          string           `json:"verdict,omitempty"`
	Stdout           string           `json:"stdout,omitempty"`
	Stderr           string           `json:"stderr,omitempty"`
	CompileOutput    string           `json:"compile_output,omitempty"`
	RuntimeMS        *int64           `json:"runtime_ms,omitempty"`
	MemoryKB         *int64           `json:"memory_kb,omitempty"`
	TestCasesPassed  int              `json:"test_cases_passed"`
	TestCasesTotal   int              `json:"test_cases_total"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
}

// CreateSubmissionRequest is the payload for submitting code.
type CreateSubmissionRequest struct {
	ProblemID  int64    `json:"problem_id"`
	Language   Language `json:"language"`
	SourceCode string   `json:"source_code"`
}

// Validate checks that required fields are present and valid.
func (r *CreateSubmissionRequest) Validate() error {
	if r.ProblemID <= 0 {
		return fmt.Errorf("problem_id must be a positive integer")
	}
	switch r.Language {
	case LanguagePython3, LanguageJava, LanguageCPP, LanguageJavaScript:
		// valid
	default:
		return fmt.Errorf("language must be one of: python3, java, cpp, javascript")
	}
	if r.SourceCode == "" {
		return fmt.Errorf("source_code is required")
	}
	if len(r.SourceCode) > 65536 {
		return fmt.Errorf("source_code exceeds maximum length of 65536 characters")
	}
	return nil
}

// CreateSubmissionResponse is returned immediately after submission creation.
type CreateSubmissionResponse struct {
	ID      int64            `json:"id"`
	Status  SubmissionStatus `json:"status"`
	Message string           `json:"message,omitempty"`
}

// IsTerminal returns true if the submission has reached a final state.
func (s *Submission) IsTerminal() bool {
	switch s.Status {
	case StatusAccepted,
		StatusWrongAnswer,
		StatusTimeLimitExceeded,
		StatusMemoryLimitExceeded,
		StatusRuntimeError,
		StatusCompilationError,
		StatusInternalError:
		return true
	}
	return false
}
