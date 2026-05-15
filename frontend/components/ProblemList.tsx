// ─────────────────────────────────────────────
//  ProblemList Component
// ─────────────────────────────────────────────

"use client";

import Link from "next/link";
import { CheckCircle2, Circle, ChevronRight } from "lucide-react";
import clsx from "clsx";
import type { ProblemSummary, Difficulty } from "@/types";

interface ProblemListProps {
  problems: ProblemSummary[];
  loading?: boolean;
  /** IDs of problems the user has already solved */
  solvedIds?: Set<number>;
}

function DifficultyBadge({ difficulty }: { difficulty: Difficulty }) {
  return (
    <span
      className={clsx(
        "inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium",
        difficulty === "easy" && "badge-easy",
        difficulty === "medium" && "badge-medium",
        difficulty === "hard" && "badge-hard"
      )}
    >
      {difficulty.charAt(0).toUpperCase() + difficulty.slice(1)}
    </span>
  );
}

function AcceptanceBar({ rate }: { rate: number }) {
  const color =
    rate >= 60 ? "bg-green-500" : rate >= 40 ? "bg-yellow-500" : "bg-red-500";

  return (
    <div className="flex items-center gap-2">
      <div className="h-1.5 w-16 overflow-hidden rounded-full bg-slate-800">
        <div
          className={clsx("h-full rounded-full transition-all", color)}
          style={{ width: `${Math.min(100, rate)}%` }}
        />
      </div>
      <span className="text-xs text-slate-400">{rate.toFixed(1)}%</span>
    </div>
  );
}

// ── Skeleton Loader ───────────────────────────

function ProblemRowSkeleton() {
  return (
    <div className="flex items-center gap-4 border-b border-slate-800/50 px-4 py-4 animate-pulse">
      <div className="h-4 w-4 rounded-full bg-slate-800" />
      <div className="h-4 w-6 rounded bg-slate-800" />
      <div className="flex-1">
        <div className="h-4 w-48 rounded bg-slate-800" />
      </div>
      <div className="h-5 w-16 rounded-full bg-slate-800" />
      <div className="hidden h-4 w-24 rounded bg-slate-800 sm:block" />
    </div>
  );
}

// ── Main Component ────────────────────────────

export default function ProblemList({
  problems,
  loading = false,
  solvedIds = new Set(),
}: ProblemListProps) {
  if (loading) {
    return (
      <div className="rounded-xl border border-slate-800 bg-slate-900 overflow-hidden">
        {/* Table header */}
        <div className="flex items-center gap-4 border-b border-slate-800 bg-slate-900/50 px-4 py-2.5">
          <div className="w-4" />
          <div className="w-8 text-xs font-medium uppercase tracking-wider text-slate-500">#</div>
          <div className="flex-1 text-xs font-medium uppercase tracking-wider text-slate-500">Title</div>
          <div className="text-xs font-medium uppercase tracking-wider text-slate-500">Difficulty</div>
          <div className="hidden text-xs font-medium uppercase tracking-wider text-slate-500 sm:block">
            Acceptance
          </div>
        </div>
        {Array.from({ length: 8 }).map((_, i) => (
          <ProblemRowSkeleton key={i} />
        ))}
      </div>
    );
  }

  if (problems.length === 0) {
    return (
      <div className="rounded-xl border border-slate-800 bg-slate-900 py-16 text-center">
        <p className="text-slate-400">No problems found.</p>
        <p className="mt-1 text-sm text-slate-500">Try adjusting your filters.</p>
      </div>
    );
  }

  return (
    <div className="rounded-xl border border-slate-800 bg-slate-900 overflow-hidden">
      {/* Table header */}
      <div className="flex items-center gap-4 border-b border-slate-800 bg-slate-900/50 px-4 py-2.5">
        <div className="w-4" />
        <div className="w-8 text-xs font-medium uppercase tracking-wider text-slate-500">#</div>
        <div className="flex-1 text-xs font-medium uppercase tracking-wider text-slate-500">Title</div>
        <div className="text-xs font-medium uppercase tracking-wider text-slate-500">Difficulty</div>
        <div className="hidden text-xs font-medium uppercase tracking-wider text-slate-500 sm:block">
          Acceptance
        </div>
        <div className="w-4" />
      </div>

      {/* Problem rows */}
      {problems.map((problem, index) => {
        const isSolved = solvedIds.has(problem.id);

        return (
          <Link
            key={problem.id}
            href={`/problems/${problem.id}`}
            className={clsx(
              "group flex items-center gap-4 border-b border-slate-800/50 px-4 py-4 transition-colors",
              "hover:bg-slate-800/40",
              index === problems.length - 1 && "border-b-0"
            )}
          >
            {/* Solved indicator */}
            <div className="w-4 flex-shrink-0">
              {isSolved ? (
                <CheckCircle2 className="h-4 w-4 text-green-400" />
              ) : (
                <Circle className="h-4 w-4 text-slate-700" />
              )}
            </div>

            {/* Problem number */}
            <div className="w-8 flex-shrink-0 font-mono text-sm text-slate-500">
              {problem.id}
            </div>

            {/* Title */}
            <div className="flex-1 min-w-0">
              <span className="font-medium text-slate-200 group-hover:text-white transition-colors truncate block">
                {problem.title}
              </span>
            </div>

            {/* Difficulty */}
            <div className="flex-shrink-0">
              <DifficultyBadge difficulty={problem.difficulty} />
            </div>

            {/* Acceptance rate */}
            <div className="hidden flex-shrink-0 sm:block">
              <AcceptanceBar rate={problem.acceptance_rate || 0} />
            </div>

            {/* Arrow */}
            <ChevronRight className="h-4 w-4 flex-shrink-0 text-slate-700 group-hover:text-slate-400 transition-colors" />
          </Link>
        );
      })}
    </div>
  );
}
