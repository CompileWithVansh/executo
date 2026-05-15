// ─────────────────────────────────────────────
//  Executo — API Client
//  Axios-based client for all backend endpoints
// ─────────────────────────────────────────────

import axios, { AxiosError, AxiosInstance, AxiosResponse } from "axios";
import type {
  Problem,
  ProblemSummary,
  Submission,
  SubmissionRequest,
  SubmissionResponse,
  PaginatedResponse,
  FilterState,
} from "@/types";

// Base URL: in production this is /api (same origin via Nginx)
// In development, Next.js rewrites /api/* to the backend
const BASE_URL =
  typeof window !== "undefined"
    ? "/api"
    : process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

// ── Axios Instance ────────────────────────────

const apiClient: AxiosInstance = axios.create({
  baseURL: BASE_URL,
  timeout: 30000,
  headers: {
    "Content-Type": "application/json",
  },
});

// Request interceptor — attach auth token if present
apiClient.interceptors.request.use(
  (config) => {
    if (typeof window !== "undefined") {
      const token = localStorage.getItem("executo_token");
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Response interceptor — normalize errors
apiClient.interceptors.response.use(
  (response: AxiosResponse) => response,
  (error: AxiosError) => {
    const message =
      (error.response?.data as { error?: string })?.error ||
      error.message ||
      "An unexpected error occurred";
    return Promise.reject(new Error(message));
  }
);

// ── Health ────────────────────────────────────

export const healthApi = {
  check: (): Promise<{ status: string; version: string }> =>
    apiClient.get("/health").then((r) => r.data),
};

// ── Problems ──────────────────────────────────

export const problemsApi = {
  /**
   * List all problems with optional filtering and pagination.
   */
  list: (
    filters?: Partial<FilterState>
  ): Promise<PaginatedResponse<ProblemSummary>> => {
    const params: Record<string, string | number> = {
      page: filters?.page || 1,
      page_size: 20,
    };
    if (filters?.difficulty && filters.difficulty !== "all") {
      params.difficulty = filters.difficulty;
    }
    if (filters?.search) {
      params.search = filters.search;
    }
    return apiClient.get("/problems", { params }).then((r) => r.data);
  },

  /**
   * Get a single problem by its numeric ID.
   */
  getById: (id: number): Promise<Problem> =>
    apiClient.get(`/problems/${id}`).then((r) => r.data),

  /**
   * Get a single problem by its slug (e.g. "two-sum").
   */
  getBySlug: (slug: string): Promise<Problem> =>
    apiClient.get(`/problems/slug/${slug}`).then((r) => r.data),
};

// ── Submissions ───────────────────────────────

export const submissionsApi = {
  /**
   * Submit code for a problem. Returns immediately with a submission ID.
   * Poll getById until status is no longer "pending" or "processing".
   */
  create: (payload: SubmissionRequest): Promise<SubmissionResponse> =>
    apiClient.post("/submissions", payload).then((r) => r.data),

  /**
   * Get a submission by ID. Use this to poll for results.
   */
  getById: (id: number): Promise<Submission> =>
    apiClient.get(`/submissions/${id}`).then((r) => r.data),

  /**
   * Get submission history for a specific problem.
   */
  listByProblem: (problemId: number): Promise<Submission[]> =>
    apiClient
      .get("/submissions", { params: { problem_id: problemId } })
      .then((r) => r.data),

  /**
   * Get all submissions for the current user.
   */
  listMine: (page = 1): Promise<PaginatedResponse<Submission>> =>
    apiClient
      .get("/submissions/me", { params: { page } })
      .then((r) => r.data),
};

// ── Polling Helper ────────────────────────────

const TERMINAL_STATUSES = new Set([
  "accepted",
  "wrong_answer",
  "time_limit_exceeded",
  "memory_limit_exceeded",
  "runtime_error",
  "compilation_error",
  "internal_error",
]);

/**
 * Poll a submission until it reaches a terminal status.
 *
 * @param submissionId - The submission ID to poll
 * @param onUpdate - Callback called on each poll with the latest submission
 * @param intervalMs - Polling interval in milliseconds (default: 2000)
 * @param maxAttempts - Maximum number of poll attempts (default: 30 = 60s)
 * @returns The final submission object
 */
export async function pollSubmission(
  submissionId: number,
  onUpdate?: (submission: Submission) => void,
  intervalMs = 2000,
  maxAttempts = 30
): Promise<Submission> {
  for (let attempt = 0; attempt < maxAttempts; attempt++) {
    const submission = await submissionsApi.getById(submissionId);
    onUpdate?.(submission);

    if (TERMINAL_STATUSES.has(submission.status)) {
      return submission;
    }

    // Wait before next poll
    await new Promise((resolve) => setTimeout(resolve, intervalMs));
  }

  // If we exhausted attempts, return the last known state
  return submissionsApi.getById(submissionId);
}

// ── Stats ─────────────────────────────────────

export interface PlatformStats {
  total_problems: number;
  total_submissions: number;
  total_accepted: number;
  acceptance_rate: number;
}

export const statsApi = {
  getPlatformStats: (): Promise<PlatformStats> =>
    apiClient.get("/stats").then((r) => r.data),
};

export default apiClient;
