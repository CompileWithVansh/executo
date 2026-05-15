-- ─────────────────────────────────────────────
--  Migration 003: Add LeetCode fields to problems
--  + create suggested_testcases table
-- ─────────────────────────────────────────────

-- Add lc_number and lc_url to problems
-- (description, examples, constraints, function_signature, test_cases
--  are kept for backward compat but will be empty for LC-linked problems)
ALTER TABLE problems
    ADD COLUMN IF NOT EXISTS lc_number  INT,
    ADD COLUMN IF NOT EXISTS lc_url     TEXT;

-- Unique index on lc_number so we can't double-insert the same problem
CREATE UNIQUE INDEX IF NOT EXISTS idx_problems_lc_number ON problems (lc_number)
    WHERE lc_number IS NOT NULL;

-- ─────────────────────────────────────────────
--  suggested_testcases
--  Users submit input + expected_output.
--  Admin approves → status = 'approved'.
--  Approved ones are shown on the problem page.
-- ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS suggested_testcases (
    id              BIGSERIAL PRIMARY KEY,
    problem_id      BIGINT NOT NULL REFERENCES problems(id) ON DELETE CASCADE,

    -- The stdin the user suggests
    input           TEXT NOT NULL,

    -- What the user believes the correct output should be
    expected_output TEXT NOT NULL,

    -- Optional note from the suggester ("handles negative numbers")
    note            TEXT,

    -- pending → approved | rejected
    status          VARCHAR(10) NOT NULL DEFAULT 'pending'
                    CHECK (status IN ('pending', 'approved', 'rejected')),

    -- Who submitted (anonymous for now — just IP or null)
    submitted_by    TEXT,

    -- Admin note on rejection
    admin_note      TEXT,

    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_stc_problem_id ON suggested_testcases (problem_id);
CREATE INDEX IF NOT EXISTS idx_stc_status     ON suggested_testcases (status);

DROP TRIGGER IF EXISTS stc_updated_at ON suggested_testcases;
CREATE TRIGGER stc_updated_at
    BEFORE UPDATE ON suggested_testcases
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
