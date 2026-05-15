// ─────────────────────────────────────────────
//  SubmissionResult Component
//  Shows verdict, runtime, memory, and output
// ─────────────────────────────────────────────

"use client";

import { CheckCircle2, XCircle, Clock, MemoryStick, AlertTriangle, Loader2, Code } from "lucide-react";
import clsx from "clsx";
import type { Submission, SubmissionStatus } from "@/types";
import { VERDICT_DISPLAY } from "@/types";

interface SubmissionResultProps {
  submission: Submission | null;
  isLoading?: boolean;
}

// ── Status Icon ───────────────────────────────

function StatusIcon({ status }: { status: SubmissionStatus }) {
  switch (status) {
    case "accepted":
      return <CheckCircle2 className="h-8 w-8 text-green-400" />;
    case "pending":
    case "processing":
      return <Loader2 className="h-8 w-8 text-blue-400 animate-spin" />;
    case "compilation_error":
    case "runtime_error":
      return <Code className="h-8 w-8 text-red-400" />;
    case "time_limit_exceeded":
      return <Clock className="h-8 w-8 text-yellow-400" />;
    case "memory_limit_exceeded":
      return <MemoryStick className="h-8 w-8 text-orange-400" />;
    default:
      return <XCircle className="h-8 w-8 text-red-400" />;
  }
}

// ── Metric Card ───────────────────────────────

function MetricCard({
  icon: Icon,
  label,
  value,
  unit,
}: {
  icon: React.ElementType;
  label: string;
  value: string | number;
  unit?: string;
}) {
  return (
    <div className="rounded-lg border border-slate-800 bg-slate-900/50 p-4">
      <div className="flex items-center gap-2 text-slate-400 mb-1">
        <Icon className="h-3.5 w-3.5" />
        <span className="text-xs uppercase tracking-wider">{label}</span>
      </div>
      <div className="text-xl font-bold text-white">
        {value}
        {unit && <span className="ml-1 text-sm font-normal text-slate-400">{unit}</span>}
      </div>
    </div>
  );
}

// ── Output Block ──────────────────────────────

function OutputBlock({
  title,
  content,
  variant = "default",
}: {
  title: string;
  content: string;
  variant?: "default" | "error" | "success";
}) {
  return (
    <div className="space-y-1.5">
      <p className="text-xs font-medium uppercase tracking-wider text-slate-500">{title}</p>
      <pre
        className={clsx(
          "rounded-lg border p-3 text-xs font-mono overflow-x-auto whitespace-pre-wrap break-words",
          variant === "error" && "border-red-800/50 bg-red-900/10 text-red-300",
          variant === "success" && "border-green-800/50 bg-green-900/10 text-green-300",
          variant === "default" && "border-slate-800 bg-slate-900 text-slate-300"
        )}
      >
        {content || "(empty)"}
      </pre>
    </div>
  );
}

// ── Main Component ────────────────────────────

export default function SubmissionResult({
  submission,
  isLoading = false,
}: SubmissionResultProps) {
  // Loading state (before submission is created)
  if (isLoading && !submission) {
    return (
      <div className="flex flex-col items-center justify-center py-16 text-center">
        <Loader2 className="mb-4 h-10 w-10 animate-spin text-blue-400" />
        <p className="font-medium text-white">Submitting your code...</p>
        <p className="mt-1 text-sm text-slate-400">Queuing for execution</p>
      </div>
    );
  }

  // No submission yet
  if (!submission) {
    return (
      <div className="flex flex-col items-center justify-center py-16 text-center">
        <div className="mb-4 rounded-full bg-slate-800 p-4">
          <Code className="h-8 w-8 text-slate-500" />
        </div>
        <p className="font-medium text-slate-300">No submission yet</p>
        <p className="mt-1 text-sm text-slate-500">
          Write your solution and click Submit to see results here.
        </p>
      </div>
    );
  }

  const verdict = VERDICT_DISPLAY[submission.status];
  const isTerminal =
    submission.status !== "pending" && submission.status !== "processing";
  const isAccepted = submission.status === "accepted";

  return (
    <div className="space-y-6 animate-fade-in">
      {/* ── Verdict Header ──────────────────── */}
      <div
        className={clsx(
          "rounded-xl border p-6 text-center",
          isAccepted
            ? "border-green-800/50 bg-green-900/10"
            : submission.status === "pending" || submission.status === "processing"
              ? "border-blue-800/50 bg-blue-900/10"
              : "border-red-800/50 bg-red-900/10"
        )}
      >
        <div className="flex justify-center mb-3">
          <StatusIcon status={submission.status} />
        </div>
        <h3 className={clsx("text-xl font-bold", verdict.color)}>
          {verdict.label}
        </h3>

        {/* Test case progress */}
        {isTerminal &&
          submission.test_cases_total !== undefined &&
          submission.test_cases_total > 0 && (
            <p className="mt-2 text-sm text-slate-400">
              {submission.test_cases_passed ?? 0} / {submission.test_cases_total} test cases passed
            </p>
          )}

        {/* Processing message */}
        {!isTerminal && (
          <p className="mt-2 text-sm text-slate-400">
            {submission.status === "pending"
              ? "Waiting in queue..."
              : "Running your code..."}
          </p>
        )}
      </div>

      {/* ── Performance Metrics ─────────────── */}
      {isTerminal && isAccepted && (
        <div className="grid grid-cols-2 gap-3">
          <MetricCard
            icon={Clock}
            label="Runtime"
            value={submission.runtime_ms ?? "—"}
            unit={submission.runtime_ms !== undefined ? "ms" : undefined}
          />
          <MetricCard
            icon={MemoryStick}
            label="Memory"
            value={
              submission.memory_kb !== undefined
                ? (submission.memory_kb / 1024).toFixed(1)
                : "—"
            }
            unit={submission.memory_kb !== undefined ? "MB" : undefined}
          />
        </div>
      )}

      {/* ── Output / Error Details ───────────── */}
      {isTerminal && (
        <div className="space-y-4">
          {/* Standard output */}
          {submission.stdout && (
            <OutputBlock
              title="Output"
              content={submission.stdout}
              variant={isAccepted ? "success" : "default"}
            />
          )}

          {/* Compilation error */}
          {submission.compile_output && (
            <OutputBlock
              title="Compilation Error"
              content={submission.compile_output}
              variant="error"
            />
          )}

          {/* Runtime error / stderr */}
          {submission.stderr && (
            <OutputBlock
              title="Error Output"
              content={submission.stderr}
              variant="error"
            />
          )}

          {/* Wrong answer hint */}
          {submission.status === "wrong_answer" && submission.verdict && (
            <OutputBlock
              title="Expected vs Got"
              content={submission.verdict}
              variant="error"
            />
          )}
        </div>
      )}

      {/* ── Submission Metadata ──────────────── */}
      {isTerminal && (
        <div className="rounded-lg border border-slate-800 bg-slate-900/30 p-3">
          <div className="flex flex-wrap gap-x-6 gap-y-1 text-xs text-slate-500">
            <span>
              Language:{" "}
              <span className="text-slate-400">
                {submission.language.toUpperCase()}
              </span>
            </span>
            <span>
              Submitted:{" "}
              <span className="text-slate-400">
                {new Date(submission.created_at).toLocaleString()}
              </span>
            </span>
            {submission.id > 0 && (
              <span>
                ID: <span className="text-slate-400 font-mono">#{submission.id}</span>
              </span>
            )}
          </div>
        </div>
      )}
    </div>
  );
}
