// ─────────────────────────────────────────────
//  Executo — Shared TypeScript Types
// ─────────────────────────────────────────────

// ── Problem ───────────────────────────────────

export type Difficulty = "easy" | "medium" | "hard";

export interface Example {
  input: string;
  output: string;
  explanation?: string;
}

export interface TestCase {
  input: string;
  expected_output: string;
}

export interface FunctionSignatures {
  python3?: string;
  java?: string;
  cpp?: string;
  javascript?: string;
  [key: string]: string | undefined;
}

export interface Problem {
  id: number;
  title: string;
  slug: string;
  description: string;
  difficulty: Difficulty;
  examples: Example[];
  constraints: string[];
  test_cases: TestCase[];
  function_signature: FunctionSignatures;
  acceptance_rate?: number;
  total_submissions?: number;
  created_at: string;
}

export interface ProblemSummary {
  id: number;
  title: string;
  slug: string;
  difficulty: Difficulty;
  acceptance_rate: number;
  total_submissions: number;
}

// ── Submission ────────────────────────────────

export type SubmissionStatus =
  | "pending"
  | "processing"
  | "accepted"
  | "wrong_answer"
  | "time_limit_exceeded"
  | "memory_limit_exceeded"
  | "runtime_error"
  | "compilation_error"
  | "internal_error";

export type Language = "python3" | "java" | "cpp" | "javascript";

export interface Submission {
  id: number;
  problem_id: number;
  language: Language;
  source_code: string;
  status: SubmissionStatus;
  verdict?: string;
  stdout?: string;
  stderr?: string;
  compile_output?: string;
  runtime_ms?: number;
  memory_kb?: number;
  test_cases_passed?: number;
  test_cases_total?: number;
  created_at: string;
  updated_at: string;
}

export interface SubmissionRequest {
  problem_id: number;
  language: Language;
  source_code: string;
}

export interface SubmissionResponse {
  id: number;
  status: SubmissionStatus;
  message?: string;
}

// ── API ───────────────────────────────────────

export interface ApiError {
  error: string;
  message?: string;
  status?: number;
}

export interface PaginatedResponse<T> {
  data: T[];
  total: number;
  page: number;
  page_size: number;
}

// ── UI State ──────────────────────────────────

export interface EditorState {
  language: Language;
  code: string;
  isSubmitting: boolean;
  isRunning: boolean;
}

export interface FilterState {
  difficulty?: Difficulty | "all";
  search?: string;
  page: number;
}

// ── Language Config ───────────────────────────

export interface LanguageConfig {
  id: Language;
  label: string;
  monacoId: string;
  judge0Id: number;
  defaultCode: string;
}

export const LANGUAGE_CONFIGS: Record<Language, LanguageConfig> = {
  python3: {
    id: "python3",
    label: "Python 3",
    monacoId: "python",
    judge0Id: 71,
    defaultCode: `# Write your solution here
class Solution:
    def solve(self):
        pass
`,
  },
  java: {
    id: "java",
    label: "Java",
    monacoId: "java",
    judge0Id: 62,
    defaultCode: `// Write your solution here
import java.util.*;

class Solution {
    public void solve() {
        // Your code here
    }
}
`,
  },
  cpp: {
    id: "cpp",
    label: "C++",
    monacoId: "cpp",
    judge0Id: 54,
    defaultCode: `// Write your solution here
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    void solve() {
        // Your code here
    }
};
`,
  },
  javascript: {
    id: "javascript",
    label: "JavaScript",
    monacoId: "javascript",
    judge0Id: 63,
    defaultCode: `// Write your solution here
/**
 * @param {*} input
 * @return {*}
 */
var solve = function(input) {
    // Your code here
};
`,
  },
};

// ── Verdict Display ───────────────────────────

export interface VerdictDisplay {
  label: string;
  color: string;
  bgColor: string;
  icon: string;
}

export const VERDICT_DISPLAY: Record<SubmissionStatus, VerdictDisplay> = {
  pending: {
    label: "Pending",
    color: "text-gray-400",
    bgColor: "bg-gray-800",
    icon: "⏳",
  },
  processing: {
    label: "Running...",
    color: "text-blue-400",
    bgColor: "bg-blue-900/30",
    icon: "⚙️",
  },
  accepted: {
    label: "Accepted",
    color: "text-green-400",
    bgColor: "bg-green-900/30",
    icon: "✓",
  },
  wrong_answer: {
    label: "Wrong Answer",
    color: "text-red-400",
    bgColor: "bg-red-900/30",
    icon: "✗",
  },
  time_limit_exceeded: {
    label: "Time Limit Exceeded",
    color: "text-yellow-400",
    bgColor: "bg-yellow-900/30",
    icon: "⏱",
  },
  memory_limit_exceeded: {
    label: "Memory Limit Exceeded",
    color: "text-orange-400",
    bgColor: "bg-orange-900/30",
    icon: "💾",
  },
  runtime_error: {
    label: "Runtime Error",
    color: "text-red-400",
    bgColor: "bg-red-900/30",
    icon: "💥",
  },
  compilation_error: {
    label: "Compilation Error",
    color: "text-red-400",
    bgColor: "bg-red-900/30",
    icon: "🔧",
  },
  internal_error: {
    label: "Internal Error",
    color: "text-gray-400",
    bgColor: "bg-gray-800",
    icon: "⚠️",
  },
};
