-- ─────────────────────────────────────────────
--  SQL Queries: Problems
--  Used with sqlc for type-safe Go code generation
-- ─────────────────────────────────────────────

-- name: GetProblemByID :one
SELECT
    id, title, slug, description, difficulty,
    examples, constraints, test_cases, function_signature,
    total_submissions, accepted_submissions,
    created_at, updated_at
FROM problems
WHERE id = $1;

-- name: GetProblemBySlug :one
SELECT
    id, title, slug, description, difficulty,
    examples, constraints, test_cases, function_signature,
    total_submissions, accepted_submissions,
    created_at, updated_at
FROM problems
WHERE slug = $1;

-- name: ListProblems :many
-- Returns paginated problem summaries with optional difficulty filter.
SELECT
    id, title, slug, difficulty,
    total_submissions,
    CASE
        WHEN total_submissions = 0 THEN 0.0
        ELSE ROUND((accepted_submissions::NUMERIC / total_submissions) * 100, 1)
    END AS acceptance_rate
FROM problems
WHERE
    ($1::VARCHAR IS NULL OR difficulty = $1)
    AND ($2::VARCHAR IS NULL OR title ILIKE '%' || $2 || '%')
ORDER BY id ASC
LIMIT $3 OFFSET $4;

-- name: CountProblems :one
SELECT COUNT(*) FROM problems
WHERE
    ($1::VARCHAR IS NULL OR difficulty = $1)
    AND ($2::VARCHAR IS NULL OR title ILIKE '%' || $2 || '%');

-- name: CreateProblem :one
INSERT INTO problems (
    title, slug, description, difficulty,
    examples, constraints, test_cases, function_signature
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id, title, slug, difficulty, created_at;

-- name: UpdateProblemStats :exec
UPDATE problems
SET
    total_submissions = total_submissions + 1,
    accepted_submissions = accepted_submissions + $2
WHERE id = $1;

-- name: GetPlatformStats :one
SELECT
    COUNT(*) AS total_problems,
    COALESCE(SUM(total_submissions), 0) AS total_submissions,
    COALESCE(SUM(accepted_submissions), 0) AS total_accepted
FROM problems;
