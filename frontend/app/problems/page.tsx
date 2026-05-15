// ─────────────────────────────────────────────
//  Problems List Page
// ─────────────────────────────────────────────

"use client";

import { useState, useEffect, useCallback } from "react";
import { Search, Filter } from "lucide-react";
import ProblemList from "@/components/ProblemList";
import { problemsApi } from "@/lib/api";
import type { ProblemSummary, Difficulty } from "@/types";

type DifficultyFilter = Difficulty | "all";

export default function ProblemsPage() {
  const [problems, setProblems] = useState<ProblemSummary[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [search, setSearch] = useState("");
  const [difficulty, setDifficulty] = useState<DifficultyFilter>("all");
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);

  const fetchProblems = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      const result = await problemsApi.list({ difficulty, search, page });
      setProblems(result.data);
      setTotal(result.total);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to load problems");
    } finally {
      setLoading(false);
    }
  }, [difficulty, search, page]);

  // Debounce search input
  useEffect(() => {
    const timer = setTimeout(() => {
      setPage(1);
      fetchProblems();
    }, 300);
    return () => clearTimeout(timer);
  }, [search, difficulty]); // eslint-disable-line react-hooks/exhaustive-deps

  useEffect(() => {
    fetchProblems();
  }, [page]); // eslint-disable-line react-hooks/exhaustive-deps

  const difficultyOptions: { value: DifficultyFilter; label: string }[] = [
    { value: "all", label: "All Difficulties" },
    { value: "easy", label: "Easy" },
    { value: "medium", label: "Medium" },
    { value: "hard", label: "Hard" },
  ];

  return (
    <div className="mx-auto max-w-5xl px-4 py-10">
      {/* ── Header ─────────────────────────────── */}
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-white">Problems</h1>
        <p className="mt-2 text-slate-400">
          {total > 0 ? `${total} problem${total !== 1 ? "s" : ""} available` : "Practice your coding skills"}
        </p>
      </div>

      {/* ── Filters ────────────────────────────── */}
      <div className="mb-6 flex flex-col gap-3 sm:flex-row">
        {/* Search */}
        <div className="relative flex-1">
          <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-slate-500" />
          <input
            type="text"
            placeholder="Search problems..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="w-full rounded-lg border border-slate-700 bg-slate-900 py-2.5 pl-10 pr-4 text-sm text-slate-200 placeholder-slate-500 outline-none transition-colors focus:border-slate-500 focus:ring-1 focus:ring-slate-500"
          />
        </div>

        {/* Difficulty filter */}
        <div className="relative">
          <Filter className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-slate-500" />
          <select
            value={difficulty}
            onChange={(e) => {
              setDifficulty(e.target.value as DifficultyFilter);
              setPage(1);
            }}
            className="appearance-none rounded-lg border border-slate-700 bg-slate-900 py-2.5 pl-10 pr-8 text-sm text-slate-200 outline-none transition-colors focus:border-slate-500 focus:ring-1 focus:ring-slate-500"
          >
            {difficultyOptions.map((opt) => (
              <option key={opt.value} value={opt.value}>
                {opt.label}
              </option>
            ))}
          </select>
        </div>
      </div>

      {/* ── Difficulty Stats ────────────────────── */}
      <div className="mb-6 grid grid-cols-3 gap-3">
        {[
          { label: "Easy", color: "text-green-400", bg: "bg-green-900/20 border-green-800/30" },
          { label: "Medium", color: "text-yellow-400", bg: "bg-yellow-900/20 border-yellow-800/30" },
          { label: "Hard", color: "text-red-400", bg: "bg-red-900/20 border-red-800/30" },
        ].map((d) => (
          <button
            key={d.label}
            onClick={() => {
              setDifficulty(d.label.toLowerCase() as DifficultyFilter);
              setPage(1);
            }}
            className={`rounded-lg border px-4 py-2.5 text-sm font-medium transition-all ${d.bg} ${d.color} hover:opacity-80`}
          >
            {d.label}
          </button>
        ))}
      </div>

      {/* ── Problem List ────────────────────────── */}
      {error ? (
        <div className="rounded-lg border border-red-800/50 bg-red-900/20 p-6 text-center">
          <p className="text-red-400">{error}</p>
          <button
            onClick={fetchProblems}
            className="mt-3 text-sm text-red-300 underline hover:text-red-200"
          >
            Try again
          </button>
        </div>
      ) : (
        <ProblemList problems={problems} loading={loading} />
      )}

      {/* ── Pagination ──────────────────────────── */}
      {!loading && total > 20 && (
        <div className="mt-8 flex items-center justify-center gap-2">
          <button
            onClick={() => setPage((p) => Math.max(1, p - 1))}
            disabled={page === 1}
            className="rounded-lg border border-slate-700 px-4 py-2 text-sm text-slate-300 transition-colors hover:border-slate-500 disabled:cursor-not-allowed disabled:opacity-40"
          >
            Previous
          </button>
          <span className="text-sm text-slate-400">
            Page {page} of {Math.ceil(total / 20)}
          </span>
          <button
            onClick={() => setPage((p) => p + 1)}
            disabled={page >= Math.ceil(total / 20)}
            className="rounded-lg border border-slate-700 px-4 py-2 text-sm text-slate-300 transition-colors hover:border-slate-500 disabled:cursor-not-allowed disabled:opacity-40"
          >
            Next
          </button>
        </div>
      )}
    </div>
  );
}
