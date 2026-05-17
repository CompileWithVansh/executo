"use client";

import { useState, useCallback, useEffect } from "react";
import dynamic from "next/dynamic";
import {
  Play, RotateCcw, ChevronDown, Terminal,
  Clock, Cpu, AlertCircle, CheckCircle2, Loader2,
} from "lucide-react";
import axios from "axios";

const Editor = dynamic(() => import("@/components/Editor"), { ssr: false });

// ── Language config ───────────────────────────
const LANGUAGES = [
  {
    id: "python3", label: "Python 3", monacoId: "python", judge0Id: 71,
    default: `# Python 3 — Executo Playground
import sys

def main():
    # Uncomment to read stdin:
    # data = sys.stdin.read().split()

    print("Hello from Executo!")
    print("Python is running 🐍")

main()
`,
  },
  {
    id: "javascript", label: "JavaScript", monacoId: "javascript", judge0Id: 63,
    default: `// JavaScript — Executo Playground

console.log("Hello from Executo!");
console.log(\`JavaScript is running 🟨\`);
`,
  },
  {
    id: "java", label: "Java", monacoId: "java", judge0Id: 62,
    default: `// Java — Executo Playground
import java.util.*;

public class Main {
    public static void main(String[] args) {
        System.out.println("Hello from Executo!");
        System.out.println("Java is running ☕");
    }
}
`,
  },
  {
    id: "cpp", label: "C++", monacoId: "cpp", judge0Id: 54,
    default: `// C++ — Executo Playground
#include <bits/stdc++.h>
using namespace std;

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);

    cout << "Hello from Executo!" << endl;
    cout << "C++ is running ⚡" << endl;

    return 0;
}
`,
  },
];

const STATUS_LABELS: Record<number, string> = {
  3: "Accepted", 4: "Wrong Answer", 5: "Time Limit Exceeded",
  6: "Compilation Error", 7: "Runtime Error (SIGSEGV)",
  8: "Runtime Error (SIGXFSZ)", 9: "Runtime Error (SIGFPE)",
  10: "Runtime Error (SIGABRT)", 11: "Runtime Error (NZEC)",
  12: "Runtime Error (Other)", 13: "Internal Error", 14: "Exec Format Error",
};

interface RunResult {
  stdout: string;
  stderr: string;
  compile_output: string;
  status: { id: number; description: string };
  time: string;
  memory: number;
}

type RunStatus = "idle" | "running" | "success" | "error";



export default function PlaygroundPage() {
  const [langIdx, setLangIdx]   = useState(0);
  const [codes, setCodes]       = useState<Record<string, string>>(
    () => Object.fromEntries(LANGUAGES.map((l) => [l.id, l.default]))
  );
  const [stdin, setStdin]       = useState("");
  const [showStdin, setShowStdin] = useState(false);
  const [result, setResult]     = useState<RunResult | null>(null);
  const [runStatus, setRunStatus] = useState<RunStatus>("idle");
  const [langOpen, setLangOpen] = useState(false);
  const [netError, setNetError] = useState<string | null>(null);

  const lang = LANGUAGES[langIdx];
  const code = codes[lang.id];

  const setCode = useCallback(
    (val: string | undefined) => setCodes((p) => ({ ...p, [lang.id]: val ?? "" })),
    [lang.id]
  );

  // Ctrl+Enter shortcut
  useEffect(() => {
    const h = () => handleRun();
    window.addEventListener("editor:submit", h);
    return () => window.removeEventListener("editor:submit", h);
  });

  const handleRun = useCallback(async () => {
    if (runStatus === "running") return;
    setRunStatus("running");
    setResult(null);
    setNetError(null);

    try {
      const submitRes = await axios.post("/api/run", {
        source_code: btoa(unescape(encodeURIComponent(code))),
        language_id: lang.judge0Id,
        stdin: stdin ? btoa(unescape(encodeURIComponent(stdin))) : "",
      });

      const token: string = submitRes.data.token;

      for (let i = 0; i < 20; i++) {
        await new Promise((r) => setTimeout(r, 1000));
        const poll = await axios.get(`/api/run/${token}`);
        const data: RunResult = poll.data;
        if (data.status.id > 2) {
          setResult(data);
          setRunStatus(data.status.id === 3 ? "success" : "error");
          return;
        }
      }

      setNetError("Execution timed out. Please try again.");
      setRunStatus("error");
    } catch (err: unknown) {
      setNetError(err instanceof Error ? err.message : "Failed to run code");
      setRunStatus("error");
    }
  }, [code, lang.judge0Id, stdin, runStatus]);

  const handleReset = () => {
    setCodes((p) => ({ ...p, [lang.id]: lang.default }));
    setResult(null);
    setRunStatus("idle");
    setNetError(null);
  };

  const stdout     = result?.stdout;
  const stderr     = result?.stderr;
  const compileErr = result?.compile_output;
  const isOk       = result?.status.id === 3;

  return (
    <div
      className="flex h-[calc(100vh-3.5rem)] flex-col"
      style={{ background: "var(--bg)" }}
    >
      {/* ── Toolbar ──────────────────────────── */}
      <div
        className="flex h-12 shrink-0 items-center justify-between border-b px-4 gap-3"
        style={{ background: "var(--surface)", borderColor: "var(--border)" }}
      >
        {/* Language selector */}
        <div className="relative">
          <button
            onClick={() => setLangOpen(!langOpen)}
            className="flex items-center gap-2 rounded-lg border px-3 py-1.5 text-sm font-medium transition-colors"
            style={{
              borderColor: "var(--border-2)",
              color: "var(--text-2)",
              background: "var(--surface-2)",
            }}
          >
            {lang.label}
            <ChevronDown className="h-3.5 w-3.5" />
          </button>

          {langOpen && (
            <div
              className="absolute left-0 top-full mt-1 z-50 w-40 rounded-xl border py-1 shadow-2xl"
              style={{ background: "var(--surface-2)", borderColor: "var(--border-2)" }}
            >
              {LANGUAGES.map((l, i) => (
                <button
                  key={l.id}
                  onClick={() => { setLangIdx(i); setLangOpen(false); }}
                  className="w-full px-3 py-2 text-left text-sm transition-colors hover:text-white"
                  style={{
                    color: i === langIdx ? "#ec4899" : "var(--text-2)",
                    fontWeight: i === langIdx ? 600 : 400,
                  }}
                >
                  {l.label}
                </button>
              ))}
            </div>
          )}
        </div>

        {/* Stdin toggle */}
        <button
          onClick={() => setShowStdin(!showStdin)}
          className="flex items-center gap-1.5 rounded-lg px-3 py-1.5 text-xs transition-colors"
          style={{ color: showStdin ? "#f59e0b" : "var(--text-3)" }}
        >
          <Terminal className="h-3.5 w-3.5" />
          {showStdin ? "Hide stdin" : "Add stdin"}
        </button>

        {/* Actions */}
        <div className="flex items-center gap-2">
          <button
            onClick={handleReset}
            className="rounded-lg p-1.5 transition-colors hover:text-white"
            style={{ color: "var(--text-3)" }}
            title="Reset to default"
          >
            <RotateCcw className="h-4 w-4" />
          </button>

          <button
            onClick={handleRun}
            disabled={runStatus === "running"}
            className="btn-primary flex items-center gap-2 rounded-lg px-4 py-1.5 text-sm"
            style={{ boxShadow: runStatus !== "running" ? "0 2px 12px rgba(236,72,153,0.3)" : "none" }}
          >
            {runStatus === "running"
              ? <Loader2 className="h-3.5 w-3.5 animate-spin" />
              : <Play className="h-3.5 w-3.5" />
            }
            {runStatus === "running" ? "Running…" : "Run"}
            <span className="hidden text-xs opacity-50 sm:inline">Ctrl+↵</span>
          </button>
        </div>
      </div>

      {/* ── Main split ───────────────────────── */}
      <div className="flex flex-1 overflow-hidden">

        {/* ── Editor pane ──────────────────── */}
        <div className="flex flex-1 flex-col overflow-hidden">
          {showStdin && (
            <div
              className="shrink-0 border-b px-4 py-3"
              style={{ borderColor: "var(--border)", background: "var(--surface)" }}
            >
              <p className="mb-1.5 text-xs font-semibold uppercase tracking-wider" style={{ color: "var(--text-3)" }}>
                Stdin
              </p>
              <textarea
                value={stdin}
                onChange={(e) => setStdin(e.target.value)}
                rows={3}
                placeholder="Enter input for your program…"
                className="w-full resize-none rounded-lg border bg-transparent p-2.5 text-xs outline-none"
                style={{
                  borderColor: "var(--border-2)",
                  color: "var(--text)",
                  fontFamily: "'JetBrains Mono', monospace",
                  caretColor: "#ec4899",
                }}
              />
            </div>
          )}

          <div className="flex-1 overflow-hidden">
            <Editor language={lang.monacoId} value={code} onChange={setCode} height="100%" />
          </div>
        </div>

        {/* ── Output pane ──────────────────── */}
        <div
          className="flex w-[380px] shrink-0 flex-col border-l"
          style={{ borderColor: "var(--border)", background: "var(--surface)" }}
        >
          {/* Output header */}
          <div
            className="flex h-10 shrink-0 items-center justify-between border-b px-4"
            style={{ borderColor: "var(--border)" }}
          >
            <span
              className="text-xs font-semibold uppercase tracking-wider"
              style={{ color: "var(--text-3)" }}
            >
              Output
            </span>
            {result && (
              <div className="flex items-center gap-3 text-xs" style={{ color: "var(--text-3)" }}>
                {result.time && (
                  <span className="flex items-center gap-1">
                    <Clock className="h-3 w-3" />
                    {result.time}s
                  </span>
                )}
                {result.memory && (
                  <span className="flex items-center gap-1">
                    <Cpu className="h-3 w-3" />
                    {(result.memory / 1024).toFixed(1)} MB
                  </span>
                )}
              </div>
            )}
          </div>

          {/* Output body */}
          <div className="flex-1 overflow-y-auto p-4 space-y-4">

            {/* Idle */}
            {runStatus === "idle" && !result && (
              <div className="flex flex-col items-center justify-center h-full text-center py-16">
                <div
                  className="mb-4 rounded-full p-4"
                  style={{ background: "var(--surface-2)" }}
                >
                  <Play className="h-5 w-5" style={{ color: "var(--text-3)" }} />
                </div>
                <p className="text-sm font-medium" style={{ color: "var(--text-2)" }}>
                  Hit Run to execute
                </p>
                <p className="mt-1 text-xs" style={{ color: "var(--text-3)" }}>
                  or press Ctrl + Enter
                </p>
              </div>
            )}

            {/* Running */}
            {runStatus === "running" && (
              <div className="flex flex-col items-center justify-center h-full text-center py-16">
                <Loader2
                  className="mb-4 h-7 w-7 animate-spin"
                  style={{ color: "#ec4899" }}
                />
                <p className="text-sm font-medium" style={{ color: "var(--text-2)" }}>
                  Executing…
                </p>
              </div>
            )}

            {/* Network error */}
            {netError && (
              <div
                className="rounded-xl border p-4 animate-fade-in"
                style={{
                  borderColor: "rgba(248,113,113,0.25)",
                  background: "rgba(248,113,113,0.06)",
                }}
              >
                <div className="flex items-center gap-2 mb-1.5">
                  <AlertCircle className="h-4 w-4" style={{ color: "var(--error)" }} />
                  <span className="text-sm font-semibold" style={{ color: "var(--error)" }}>Error</span>
                </div>
                <p className="text-xs" style={{ color: "var(--text-2)" }}>{netError}</p>
              </div>
            )}

            {/* Result */}
            {result && !netError && (
              <div className="animate-fade-in space-y-4">

                {/* Status */}
                <div
                  className="flex items-center gap-2.5 rounded-xl border p-3"
                  style={{
                    borderColor: isOk ? "rgba(74,222,128,0.25)" : "rgba(248,113,113,0.25)",
                    background:  isOk ? "rgba(74,222,128,0.06)"  : "rgba(248,113,113,0.06)",
                  }}
                >
                  {isOk
                    ? <CheckCircle2 className="h-4 w-4 shrink-0" style={{ color: "var(--success)" }} />
                    : <AlertCircle  className="h-4 w-4 shrink-0" style={{ color: "var(--error)"   }} />
                  }
                  <span
                    className="text-sm font-semibold"
                    style={{ color: isOk ? "var(--success)" : "var(--error)" }}
                  >
                    {STATUS_LABELS[result.status.id] ?? result.status.description}
                  </span>
                </div>

                {/* stdout */}
                {stdout && (
                  <div>
                    <p className="mb-1.5 text-xs font-semibold uppercase tracking-wider" style={{ color: "var(--text-3)" }}>
                      Output
                    </p>
                    <pre
                      className="rounded-xl border p-3 text-xs overflow-x-auto whitespace-pre-wrap break-words"
                      style={{
                        borderColor: "var(--border)",
                        background: "var(--surface-2)",
                        color: "var(--text)",
                        fontFamily: "'JetBrains Mono', monospace",
                      }}
                    >
                      {stdout}
                    </pre>
                  </div>
                )}

                {/* Compile error */}
                {compileErr && (
                  <div>
                    <p className="mb-1.5 text-xs font-semibold uppercase tracking-wider" style={{ color: "var(--error)" }}>
                      Compilation Error
                    </p>
                    <pre
                      className="rounded-xl border p-3 text-xs overflow-x-auto whitespace-pre-wrap break-words"
                      style={{
                        borderColor: "rgba(248,113,113,0.2)",
                        background: "rgba(248,113,113,0.05)",
                        color: "#fca5a5",
                        fontFamily: "'JetBrains Mono', monospace",
                      }}
                    >
                      {compileErr}
                    </pre>
                  </div>
                )}

                {/* stderr */}
                {stderr && (
                  <div>
                    <p className="mb-1.5 text-xs font-semibold uppercase tracking-wider" style={{ color: "var(--error)" }}>
                      Stderr
                    </p>
                    <pre
                      className="rounded-xl border p-3 text-xs overflow-x-auto whitespace-pre-wrap break-words"
                      style={{
                        borderColor: "rgba(248,113,113,0.2)",
                        background: "rgba(248,113,113,0.05)",
                        color: "#fca5a5",
                        fontFamily: "'JetBrains Mono', monospace",
                      }}
                    >
                      {stderr}
                    </pre>
                  </div>
                )}

                {/* No output */}
                {!stdout && !stderr && !compileErr && isOk && (
                  <p className="text-xs" style={{ color: "var(--text-3)" }}>
                    Program exited successfully with no output.
                  </p>
                )}
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
