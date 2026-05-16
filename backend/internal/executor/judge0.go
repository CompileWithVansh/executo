// Package executor provides a client for the Judge0 CE API.
// Judge0 is a self-hosted code execution system that runs code in isolated Docker containers.
package executor

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Judge0Client is a client for the Judge0 CE REST API.
type Judge0Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// NewJudge0Client creates a new Judge0 API client.
// It reads JUDGE0_URL and JUDGE0_API_KEY from environment variables.
func NewJudge0Client() *Judge0Client {
	baseURL := os.Getenv("JUDGE0_URL")
	if baseURL == "" {
		baseURL = "http://localhost:2358"
	}
	// Remove trailing slash
	baseURL = strings.TrimRight(baseURL, "/")

	return &Judge0Client{
		baseURL: baseURL,
		apiKey:  os.Getenv("JUDGE0_API_KEY"),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// ── Judge0 API Types ──────────────────────────

// SubmitRequest is the payload sent to Judge0 to create a submission.
type SubmitRequest struct {
	SourceCode     string `json:"source_code"`      // base64-encoded
	LanguageID     int    `json:"language_id"`
	Stdin          string `json:"stdin,omitempty"`  // base64-encoded
	ExpectedOutput string `json:"expected_output,omitempty"` // base64-encoded
	CPUTimeLimit   float64 `json:"cpu_time_limit,omitempty"`
	MemoryLimit    int    `json:"memory_limit,omitempty"`
}

// SubmitResponse is returned by Judge0 after creating a submission.
type SubmitResponse struct {
	Token string `json:"token"`
}

// Judge0Status represents the execution status from Judge0.
type Judge0Status struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// Judge0Result is the full result returned by Judge0 after execution.
type Judge0Result struct {
	Token         string       `json:"token"`
	Stdout        string       `json:"stdout"`        // base64-encoded
	Stderr        string       `json:"stderr"`        // base64-encoded
	CompileOutput string       `json:"compile_output"` // base64-encoded
	Message       string       `json:"message"`
	Status        Judge0Status `json:"status"`
	Time          string       `json:"time"`   // seconds as string, e.g. "0.042"
	Memory        int          `json:"memory"` // kilobytes
}

// Judge0 status IDs
const (
	StatusInQueue     = 1
	StatusProcessing  = 2
	StatusAccepted    = 3
	StatusWrongAnswer = 4
	StatusTLE         = 5
	StatusCompileErr  = 6
	StatusRuntimeErr  = 7 // SIGSEGV
	StatusRuntimeErr2 = 8 // SIGXFSZ
	StatusRuntimeErr3 = 9 // SIGFPE
	StatusRuntimeErr4 = 10 // SIGABRT
	StatusRuntimeErr5 = 11 // NZEC
	StatusRuntimeErr6 = 12 // Other
	StatusInternalErr = 13
	StatusExecFormat  = 14
)

// ── Submit ────────────────────────────────────

// Submit sends code to Judge0 for execution and returns a token.
// The token is used to poll for results.
func (c *Judge0Client) Submit(
	sourceCode string,
	languageID int,
	stdin string,
	expectedOutput string,
) (string, error) {
	req := SubmitRequest{
		SourceCode:     base64.StdEncoding.EncodeToString([]byte(sourceCode)),
		LanguageID:     languageID,
		CPUTimeLimit:   5.0,  // 5 second CPU time limit
		MemoryLimit:    262144, // 256 MB memory limit
	}

	if stdin != "" {
		req.Stdin = base64.StdEncoding.EncodeToString([]byte(stdin))
	}
	if expectedOutput != "" {
		req.ExpectedOutput = base64.StdEncoding.EncodeToString([]byte(expectedOutput))
	}

	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshaling submit request: %w", err)
	}

	httpReq, err := http.NewRequest(
		http.MethodPost,
		c.baseURL+"/submissions?base64_encoded=true&wait=false",
		bytes.NewReader(body),
	)
	if err != nil {
		return "", fmt.Errorf("creating HTTP request: %w", err)
	}

	c.setHeaders(httpReq)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("sending submit request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("judge0 returned status %d: %s", resp.StatusCode, string(respBody))
	}

	var submitResp SubmitResponse
	if err := json.NewDecoder(resp.Body).Decode(&submitResp); err != nil {
		return "", fmt.Errorf("decoding submit response: %w", err)
	}

	return submitResp.Token, nil
}

// ── SubmitRaw ─────────────────────────────────

// SubmitRaw sends already-base64-encoded code to Judge0 (used by the playground).
// Unlike Submit(), this does NOT re-encode the source code.
func (c *Judge0Client) SubmitRaw(sourceCodeB64 string, languageID int, stdinB64 string) (string, error) {
	req := SubmitRequest{
		SourceCode:   sourceCodeB64,
		LanguageID:   languageID,
		CPUTimeLimit: 5.0,
		MemoryLimit:  262144,
	}

	if stdinB64 != "" {
		req.Stdin = stdinB64
	}

	body, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshaling submit request: %w", err)
	}

	httpReq, err := http.NewRequest(
		http.MethodPost,
		c.baseURL+"/submissions?base64_encoded=true&wait=false",
		bytes.NewReader(body),
	)
	if err != nil {
		return "", fmt.Errorf("creating HTTP request: %w", err)
	}

	c.setHeaders(httpReq)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("sending submit request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("judge0 returned status %d: %s", resp.StatusCode, string(respBody))
	}

	var submitResp SubmitResponse
	if err := json.NewDecoder(resp.Body).Decode(&submitResp); err != nil {
		return "", fmt.Errorf("decoding submit response: %w", err)
	}

	return submitResp.Token, nil
}

// ── Poll ──────────────────────────────────────

// GetResult fetches the current result for a submission token.
func (c *Judge0Client) GetResult(token string) (*Judge0Result, error) {
	url := fmt.Sprintf("%s/submissions/%s?base64_encoded=true", c.baseURL, token)

	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating HTTP request: %w", err)
	}

	c.setHeaders(httpReq)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("fetching result: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("judge0 returned status %d: %s", resp.StatusCode, string(respBody))
	}

	var result Judge0Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding result: %w", err)
	}

	// Decode base64 fields
	result.Stdout = decodeBase64(result.Stdout)
	result.Stderr = decodeBase64(result.Stderr)
	result.CompileOutput = decodeBase64(result.CompileOutput)

	return &result, nil
}

// PollUntilDone polls Judge0 until the submission reaches a terminal state.
// It uses exponential backoff starting at 500ms.
func (c *Judge0Client) PollUntilDone(token string, maxAttempts int) (*Judge0Result, error) {
	backoff := 500 * time.Millisecond

	for attempt := 0; attempt < maxAttempts; attempt++ {
		result, err := c.GetResult(token)
		if err != nil {
			return nil, err
		}

		// Status IDs 1 and 2 mean still running
		if result.Status.ID > StatusProcessing {
			return result, nil
		}

		log.Printf("Judge0 submission %s: status=%s (attempt %d/%d)",
			token, result.Status.Description, attempt+1, maxAttempts)

		time.Sleep(backoff)
		if backoff < 3*time.Second {
			backoff = time.Duration(float64(backoff) * 1.5)
		}
	}

	return nil, fmt.Errorf("submission %s did not complete after %d attempts", token, maxAttempts)
}

// ── Helpers ───────────────────────────────────

// setHeaders adds required headers to a Judge0 API request.
func (c *Judge0Client) setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// RapidAPI headers (only needed when using the hosted Judge0 service)
	if c.apiKey != "" {
		req.Header.Set("X-RapidAPI-Key", c.apiKey)
		req.Header.Set("X-RapidAPI-Host", "judge0-ce.p.rapidapi.com")
	}
}

// decodeBase64 decodes a base64 string, returning the original on error.
func decodeBase64(s string) string {
	if s == "" {
		return ""
	}
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		// Try standard decoding without padding
		decoded, err = base64.RawStdEncoding.DecodeString(s)
		if err != nil {
			return s // Return as-is if decoding fails
		}
	}
	return string(decoded)
}

// IsTerminalStatus returns true if the Judge0 status ID indicates completion.
func IsTerminalStatus(statusID int) bool {
	return statusID > StatusProcessing
}
