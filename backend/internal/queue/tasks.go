// Package queue defines Asynq task types and payloads for async job processing.
package queue

import (
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

// Task type constants — used as Asynq task type identifiers.
const (
	// TypeExecuteSubmission is the task type for processing a code submission.
	TypeExecuteSubmission = "submission:execute"
)

// Queue names
const (
	QueueCritical = "critical"
	QueueDefault  = "default"
	QueueLow      = "low"
)

// ── Execute Submission Task ───────────────────

// ExecuteSubmissionPayload is the data carried by a submission execution task.
type ExecuteSubmissionPayload struct {
	// SubmissionID is the database ID of the submission to process.
	SubmissionID int64 `json:"submission_id"`
}

// NewExecuteSubmissionTask creates a new Asynq task for executing a submission.
// The task is placed in the default queue.
func NewExecuteSubmissionTask(submissionID int64) (*asynq.Task, error) {
	payload, err := json.Marshal(ExecuteSubmissionPayload{
		SubmissionID: submissionID,
	})
	if err != nil {
		return nil, fmt.Errorf("marshaling task payload: %w", err)
	}

	return asynq.NewTask(
		TypeExecuteSubmission,
		payload,
		// Retry up to 3 times on failure
		asynq.MaxRetry(3),
		// Place in default queue
		asynq.Queue(QueueDefault),
	), nil
}

// ParseExecuteSubmissionPayload extracts the payload from an Asynq task.
func ParseExecuteSubmissionPayload(task *asynq.Task) (*ExecuteSubmissionPayload, error) {
	var payload ExecuteSubmissionPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return nil, fmt.Errorf("parsing task payload: %w", err)
	}
	return &payload, nil
}
