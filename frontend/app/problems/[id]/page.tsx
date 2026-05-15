// ─────────────────────────────────────────────
//  Problem Detail Page — Editor + Submission
// ─────────────────────────────────────────────

"use client";

import { useState, useEffect, useCallback } from "react";
import { useParams } from "next/navigation";
import { ChevronLeft, Play, Send, Clock, MemoryStick, CheckCircle2, XCircle, AlertCircle, Loader2 } from "lucide-react";
import Link from "next/link";
import Editor from "@/components/Editor";
import SubmissionResult from "@/components/SubmissionResult";
import { problemsApi, submissionsApi, pollSubmission } from "@/lib/api";
import type { Problem, Language, Submission, SubmissionStatus } from "@/types";
import { LANGUAGE_CONFIGS } from "@/types";

type TabId = "description" | "submissions" | "result";

export default function ProblemDetailPage() {
  const params = useParams();
  const problemId = Number(params.id);

  // ── State ─────────────────────────────────
  const [problem, setProblem] = useState<Problem | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [language, setLanguage] = useState<Language>("python3");
  const [code, setCode] = useState<string>("");

  const [activeTab, setActiveTab] = useState<TabId>("description");
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [currentSubmission, setCurrentSubmission] = useState<Submission | null>(null);
  const [submissionHistory, setSubmissionHistory] = useState<Submission[]>([]);

  // ── Load Problem ──────────────────────────
  useEffect(() => {
    async function loadProblem() {
      setLoading(true);
      setError(null);
      try {
        const p = await problemsApi.getById(problemId);
        setProblem(p);
        // Set default code from function signature
        const sig = p.function_signature?.[language];
        setCode(sig || LANGUAGE_CONFIGS[language].defaultCode);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Failed to load problem");
      } finally {
        setLoading(false);
      }
    }
    if (problemId) loadProblem();
  }, [problemId]);

  // Update code template when language changes
  useEffect(() => {
    if (problem) {
      const sig = problem.function_signature?.[language];
      setCode(sig || LANGUAGE_CONFIGS[language].defaultCode);
    }
  }, [language]); // eslint-disable-line react-hooks/exhaustive-deps

  // ── Submit Code ───────────────────────────
  const handleSubmit = useCallback(async () => {
    if (!problem || isSubmitting) return;

    setIsSubmitting(true);
    setActiveTab("result");
    setCurrentSubmission(null);

    try {
      // Create submission
      const response = await submissionsApi.create({
        problem_id: problem.id,
        language,
        source_code: code,
      });

      // Poll until terminal status
      const finalSubmission = await pollSubmission(
        response.id,
        (submission) => setCurrentSubmission(submission)
      );

      setCurrentSubmission(finalSubmission);
      setSubmissionHistory((prev) => [finalSubmission, ...prev]);
    } catch (err) {
      const errorSubmission: Submission = {
        id: 0,
        problem_id: problemId,
        language,
        source_code: code,
        status: "internal_error",
        stderr: err instanceof Error ? err.message : "Submission failed",
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString(),
      };
      setCurrentSubmission(errorSubmission);
    } finally {
      setIsSubmitting(false);
    }
  }, [problem, isSubmitting, language, code, problemId]);

  // ── Render ────────────────────────────────

  if (loading) {
    return (
      <div className="flex h-[calc(100vh-4rem)] items-center justify-center">
        <div className="flex flex-col items-center gap-3">
          <Loader2 className="h-8 w-8 animate-spin text-green-400" />
          <p className="text-slate-400">Loading problem...</p>
        </div>
      </div>
    );
  }

  if (error || !problem) {
    return (
      <div className="flex h-[calc(100vh-4rem)] items-center justify-center">
        <div className="text-center">
          <XCircle className="mx-auto mb-3 h-12 w-12 text-red-400" />
          <p className="text-lg font-semibold text-white">Problem not found</p>
          <p className="mt-1 text-slate-400">{error}</p>
          <Link href="/problems" className="mt-4 inline-block text-green-400 hover:text-green-300">
            ← Back to problems
          </Link>
        </div>
      </div>
    );
  }

  const difficultyColor =
    problem.difficulty === "easy"
      ? "text-green-400"
      : problem.difficulty === "medium"
        ? "text-yellow-400"
        : "text-red-400";

  return (
    <div className="flex h-[calc(100vh-4rem)] flex-col overflow-hidden lg:flex-row">
      {/* ── Left Panel: Problem Description ──── */}
      <div className="flex w-full flex-col border-b border-slate-800 lg:w-[45%] lg:border-b-0 lg:border-r">
        {/* Problem header */}
        <div className="border-b border-slate-800 px-4 py-3">
          <Link
            href="/problems"
            className="mb-2 flex items-center gap-1 text-xs text-slate-500 hover:text-slate-300"
          >
            <ChevronLeft className="h-3 w-3" />
            Problems
          </Link>
          <div className="flex items-center gap-3">
            <h1 className="text-lg font-bold text-white">{problem.title}</h1>
            <span className={`text-sm font-medium ${difficultyColor}`}>
              {problem.difficulty.charAt(0).toUpperCase() + problem.difficulty.slice(1)}
            </span>
          </div>
        </div>

        {/* Tabs */}
        <div className="flex border-b border-slate-800">
          {(["description", "submissions", "result"] as TabId[]).map((tab) => (
            <button
              key={tab}
              onClick={() => setActiveTab(tab)}
              className={`px-4 py-2.5 text-sm font-medium capitalize transition-colors ${
                activeTab === tab
                  ? "border-b-2 border-green-400 text-green-400"
                  : "text-slate-400 hover:text-slate-200"
              }`}
            >
              {tab}
              {tab === "result" && currentSubmission && (
                <span
                  className={`ml-2 inline-block h-2 w-2 rounded-full ${
                    currentSubmission.status === "accepted"
                      ? "bg-green-400"
                      : currentSubmission.status === "pending" || currentSubmission.status === "processing"
                        ? "bg-blue-400 animate-pulse"
                        : "bg-red-400"
                  }`}
                />
              )}
            </button>
          ))}
        </div>

        {/* Tab content */}
        <div className="flex-1 overflow-y-auto p-4">
          {activeTab === "description" && (
            <div className="problem-description space-y-6">
              {/* Description */}
              <div>
                <p className="text-slate-300 leading-relaxed">{problem.description}</p>
              </div>

              {/* Examples */}
              {problem.examples?.length > 0 && (
                <div>
                  {problem.examples.map((example, i) => (
                    <div key={i} className="mb-4">
                      <p className="mb-2 font-semibold text-white">Example {i + 1}:</p>
                      <div className="rounded-lg bg-slate-900 border border-slate-800 p-4 space-y-2">
                        <div>
                          <span className="text-slate-400 text-sm">Input: </span>
                          <code className="text-slate-200 text-sm">{example.input}</code>
                        </div>
                        <div>
                          <span className="text-slate-400 text-sm">Output: </span>
                          <code className="text-slate-200 text-sm">{example.output}</code>
                        </div>
                        {example.explanation && (
                          <div>
                            <span className="text-slate-400 text-sm">Explanation: </span>
                            <span className="text-slate-300 text-sm">{example.explanation}</span>
                          </div>
                        )}
                      </div>
                    </div>
                  ))}
                </div>
              )}

              {/* Constraints */}
              {problem.constraints?.length > 0 && (
                <div>
                  <p className="mb-2 font-semibold text-white">Constraints:</p>
                  <ul className="space-y-1">
                    {problem.constraints.map((c, i) => (
                      <li key={i} className="flex items-start gap-2 text-sm text-slate-300">
                        <span className="mt-1.5 h-1.5 w-1.5 flex-shrink-0 rounded-full bg-slate-500" />
                        <code>{c}</code>
                      </li>
                    ))}
                  </ul>
                </div>
              )}
            </div>
          )}

          {activeTab === "submissions" && (
            <div>
              {submissionHistory.length === 0 ? (
                <div className="py-12 text-center text-slate-500">
                  <p>No submissions yet.</p>
                  <p className="mt-1 text-sm">Submit your solution to see results here.</p>
                </div>
              ) : (
                <div className="space-y-3">
                  {submissionHistory.map((sub, i) => (
                    <div
                      key={sub.id || i}
                      className="rounded-lg border border-slate-800 bg-slate-900 p-4"
                    >
                      <div className="flex items-center justify-between">
                        <span
                          className={`text-sm font-medium ${
                            sub.status === "accepted" ? "text-green-400" : "text-red-400"
                          }`}
                        >
                          {sub.status.replace(/_/g, " ").replace(/\b\w/g, (c) => c.toUpperCase())}
                        </span>
                        <span className="text-xs text-slate-500">
                          {new Date(sub.created_at).toLocaleString()}
                        </span>
                      </div>
                      {sub.runtime_ms !== undefined && (
                        <div className="mt-2 flex gap-4 text-xs text-slate-400">
                          <span>Runtime: {sub.runtime_ms}ms</span>
                          {sub.memory_kb && <span>Memory: {sub.memory_kb}KB</span>}
                        </div>
                      )}
                    </div>
                  ))}
                </div>
              )}
            </div>
          )}

          {activeTab === "result" && (
            <SubmissionResult submission={currentSubmission} isLoading={isSubmitting} />
          )}
        </div>
      </div>

      {/* ── Right Panel: Editor ───────────────── */}
      <div className="flex flex-1 flex-col">
        {/* Editor toolbar */}
        <div className="flex items-center justify-between border-b border-slate-800 px-4 py-2">
          {/* Language selector */}
          <select
            value={language}
            onChange={(e) => setLanguage(e.target.value as Language)}
            className="rounded-md border border-slate-700 bg-slate-900 px-3 py-1.5 text-sm text-slate-200 outline-none focus:border-slate-500"
          >
            {Object.values(LANGUAGE_CONFIGS).map((lang) => (
              <option key={lang.id} value={lang.id}>
                {lang.label}
              </option>
            ))}
          </select>

          {/* Action buttons */}
          <div className="flex items-center gap-2">
            <button
              onClick={handleSubmit}
              disabled={isSubmitting}
              className="flex items-center gap-2 rounded-lg bg-green-600 px-4 py-1.5 text-sm font-medium text-white transition-colors hover:bg-green-500 disabled:cursor-not-allowed disabled:opacity-60"
            >
              {isSubmitting ? (
                <>
                  <Loader2 className="h-3.5 w-3.5 animate-spin" />
                  Running...
                </>
              ) : (
                <>
                  <Send className="h-3.5 w-3.5" />
                  Submit
                </>
              )}
            </button>
          </div>
        </div>

        {/* Monaco Editor */}
        <div className="flex-1 overflow-hidden">
          <Editor
            language={LANGUAGE_CONFIGS[language].monacoId}
            value={code}
            onChange={(val) => setCode(val || "")}
          />
        </div>

        {/* Status bar */}
        <div className="flex items-center justify-between border-t border-slate-800 px-4 py-1.5 text-xs text-slate-500">
          <span>{LANGUAGE_CONFIGS[language].label}</span>
          {currentSubmission && currentSubmission.status !== "pending" && currentSubmission.status !== "processing" && (
            <span
              className={
                currentSubmission.status === "accepted" ? "text-green-400" : "text-red-400"
              }
            >
              Last: {currentSubmission.status.replace(/_/g, " ")}
              {currentSubmission.runtime_ms !== undefined &&
                ` · ${currentSubmission.runtime_ms}ms`}
            </span>
          )}
        </div>
      </div>
    </div>
  );
}
