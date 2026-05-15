-- ─────────────────────────────────────────────
--  Migration 001: Create problems table
-- ─────────────────────────────────────────────

CREATE TABLE IF NOT EXISTS problems (
    id                  BIGSERIAL PRIMARY KEY,
    title               VARCHAR(255) NOT NULL,
    slug                VARCHAR(255) NOT NULL UNIQUE,
    description         TEXT NOT NULL,
    difficulty          VARCHAR(10) NOT NULL CHECK (difficulty IN ('easy', 'medium', 'hard')),

    -- JSON arrays stored as TEXT for broad compatibility
    -- examples: [{"input": "...", "output": "...", "explanation": "..."}]
    examples            TEXT NOT NULL DEFAULT '[]',

    -- constraints: ["2 <= nums.length <= 10^4", ...]
    constraints         TEXT NOT NULL DEFAULT '[]',

    -- test_cases: [{"input": "...", "expected_output": "..."}]
    -- These are hidden from users; used by the judge
    test_cases          TEXT NOT NULL DEFAULT '[]',

    -- function_signature: {"python3": "def twoSum...", "java": "public int[]..."}
    function_signature  TEXT NOT NULL DEFAULT '{}',

    -- Computed stats (updated by triggers or background jobs)
    total_submissions   BIGINT NOT NULL DEFAULT 0,
    accepted_submissions BIGINT NOT NULL DEFAULT 0,

    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Index for slug lookups (used in URL routing)
CREATE INDEX IF NOT EXISTS idx_problems_slug ON problems (slug);

-- Index for difficulty filtering
CREATE INDEX IF NOT EXISTS idx_problems_difficulty ON problems (difficulty);

-- Trigger to auto-update updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS problems_updated_at ON problems;
CREATE TRIGGER problems_updated_at
    BEFORE UPDATE ON problems
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
