import Link from "next/link";
import { Terminal, Zap, Globe, Lock, ArrowRight, Play } from "lucide-react";

const LANGUAGES = [
  { name: "Python 3",   hex: "#3b82f6", icon: "🐍" },
  { name: "Java",       hex: "#f97316", icon: "☕" },
  { name: "C++",        hex: "#a855f7", icon: "⚡" },
  { name: "JavaScript", hex: "#eab308", icon: "🟨" },
];

const FEATURES = [
  {
    icon: Zap,
    title: "Instant Execution",
    desc: "Code runs in isolated containers via Judge0. Results in under 3 seconds.",
    iconColor: "#ec4899",
    iconBg: "rgba(236,72,153,0.1)",
  },
  {
    icon: Globe,
    title: "4 Languages",
    desc: "Python 3, Java, C++, and JavaScript — full standard library support.",
    iconColor: "#f59e0b",
    iconBg: "rgba(245,158,11,0.1)",
  },
  {
    icon: Lock,
    title: "Sandboxed & Safe",
    desc: "Every run is isolated with strict CPU, memory, and time limits.",
    iconColor: "#ec4899",
    iconBg: "rgba(236,72,153,0.1)",
  },
  {
    icon: Terminal,
    title: "VS Code Editor",
    desc: "Monaco editor with syntax highlighting, autocomplete, and bracket matching.",
    iconColor: "#f59e0b",
    iconBg: "rgba(245,158,11,0.1)",
  },
];

export default function HomePage() {
  return (
    <div className="flex flex-col">

      {/* ── Hero ─────────────────────────────── */}
      <section
        className="relative overflow-hidden border-b"
        style={{ borderColor: "var(--border)" }}
      >
        {/* Very subtle warm grid */}
        <div
          className="absolute inset-0 pointer-events-none"
          style={{
            backgroundImage:
              "linear-gradient(rgba(236,72,153,0.04) 1px, transparent 1px), linear-gradient(90deg, rgba(236,72,153,0.04) 1px, transparent 1px)",
            backgroundSize: "52px 52px",
          }}
        />
        {/* Glow — rose left, gold right, very dim */}
        <div
          className="absolute -top-32 left-1/4 w-[480px] h-[320px] rounded-full blur-[140px] pointer-events-none"
          style={{ background: "rgba(236,72,153,0.07)" }}
        />
        <div
          className="absolute -top-32 right-1/4 w-[480px] h-[320px] rounded-full blur-[140px] pointer-events-none"
          style={{ background: "rgba(245,158,11,0.06)" }}
        />

        <div className="relative mx-auto max-w-5xl px-5 py-28 text-center">
          {/* Badge */}
          <div
            className="mb-7 inline-flex items-center gap-2 rounded-full border px-4 py-1.5 text-xs font-medium"
            style={{
              borderColor: "rgba(236,72,153,0.25)",
              background: "rgba(236,72,153,0.06)",
              color: "#ec4899",
            }}
          >
            <span
              className="h-1.5 w-1.5 rounded-full animate-pulse-dot"
              style={{ background: "#ec4899" }}
            />
            No signup · No install · Just code
          </div>

          <h1
            className="mb-5 text-5xl font-bold tracking-tight sm:text-6xl"
            style={{ color: "var(--text)", lineHeight: 1.1 }}
          >
            Write code.{" "}
            <span className="gradient-text">Run it instantly.</span>
          </h1>

          <p
            className="mx-auto mb-10 max-w-lg text-base leading-relaxed"
            style={{ color: "var(--text-2)" }}
          >
            Executo is a fast online compiler for Python, Java, C++, and JavaScript.
            Open the playground, write your code, and see output in seconds.
          </p>

          <div className="flex flex-col items-center gap-3 sm:flex-row sm:justify-center">
            <Link
              href="/playground"
              className="btn-primary inline-flex items-center gap-2 rounded-xl px-7 py-3 text-sm shadow-lg"
              style={{ boxShadow: "0 4px 24px rgba(236,72,153,0.25)" }}
            >
              <Play className="h-4 w-4" />
              Open Playground
            </Link>
            <Link
              href="/problems"
              className="btn-ghost inline-flex items-center gap-2 rounded-xl px-7 py-3 text-sm font-semibold"
            >
              Browse Problems
              <ArrowRight className="h-4 w-4" />
            </Link>
          </div>

          {/* Language pills */}
          <div className="mt-12 flex flex-wrap justify-center gap-2.5">
            {LANGUAGES.map((lang) => (
              <span
                key={lang.name}
                className="flex items-center gap-1.5 rounded-full border px-4 py-1.5 text-xs font-medium"
                style={{
                  borderColor: `${lang.hex}28`,
                  background: `${lang.hex}0d`,
                  color: lang.hex,
                }}
              >
                {lang.icon} {lang.name}
              </span>
            ))}
          </div>
        </div>
      </section>

      {/* ── Stats ────────────────────────────── */}
      <section
        className="border-b"
        style={{ borderColor: "var(--border)", background: "var(--surface)" }}
      >
        <div className="mx-auto max-w-5xl px-5 py-8">
          <div className="grid grid-cols-2 gap-6 sm:grid-cols-4 text-center">
            {[
              { label: "Languages",   value: "4"     },
              { label: "Avg Runtime", value: "< 2s"  },
              { label: "Uptime",      value: "99.9%" },
              { label: "Cost",        value: "Free"  },
            ].map((s) => (
              <div key={s.label}>
                <div className="text-2xl font-bold gradient-text">{s.value}</div>
                <div className="mt-1 text-xs" style={{ color: "var(--text-3)" }}>{s.label}</div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* ── Features ─────────────────────────── */}
      <section className="mx-auto w-full max-w-5xl px-5 py-20">
        <div className="mb-12 text-center">
          <h2 className="text-2xl font-bold" style={{ color: "var(--text)" }}>
            Everything you need to run code
          </h2>
          <p className="mt-2 text-sm" style={{ color: "var(--text-2)" }}>
            No environment setup. No dependencies. Open and code.
          </p>
        </div>

        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
          {FEATURES.map((f) => (
            <div
              key={f.title}
              className="card-hover rounded-xl border p-5"
              style={{ background: "var(--surface)", borderColor: "var(--border)" }}
            >
              <div
                className="mb-4 inline-flex rounded-lg p-2.5"
                style={{ background: f.iconBg }}
              >
                <f.icon className="h-4 w-4" style={{ color: f.iconColor }} />
              </div>
              <h3 className="mb-1.5 text-sm font-semibold" style={{ color: "var(--text)" }}>
                {f.title}
              </h3>
              <p className="text-xs leading-relaxed" style={{ color: "var(--text-2)" }}>
                {f.desc}
              </p>
            </div>
          ))}
        </div>
      </section>

      {/* ── How it works ─────────────────────── */}
      <section
        className="border-t border-b"
        style={{ borderColor: "var(--border)", background: "var(--surface)" }}
      >
        <div className="mx-auto max-w-5xl px-5 py-20">
          <div className="mb-12 text-center">
            <h2 className="text-2xl font-bold" style={{ color: "var(--text)" }}>
              How it works
            </h2>
          </div>
          <div className="grid gap-10 sm:grid-cols-3">
            {[
              { n: "01", title: "Write",      desc: "Open the playground and write your code in the Monaco editor." },
              { n: "02", title: "Run",         desc: "Hit Run. Your code executes in an isolated Docker container via Judge0." },
              { n: "03", title: "See output",  desc: "stdout, stderr, runtime, and memory — back in under 3 seconds." },
            ].map((s) => (
              <div key={s.n}>
                <div
                  className="mb-3 text-5xl font-black"
                  style={{
                    background: "linear-gradient(135deg, #ec4899, #f59e0b)",
                    WebkitBackgroundClip: "text",
                    WebkitTextFillColor: "transparent",
                    opacity: 0.25,
                  }}
                >
                  {s.n}
                </div>
                <h3 className="mb-1.5 font-semibold" style={{ color: "var(--text)" }}>{s.title}</h3>
                <p className="text-sm leading-relaxed" style={{ color: "var(--text-2)" }}>{s.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* ── CTA ──────────────────────────────── */}
      <section className="mx-auto max-w-2xl px-5 py-24 text-center">
        <h2 className="mb-4 text-3xl font-bold" style={{ color: "var(--text)" }}>
          Ready to run some code?
        </h2>
        <p className="mb-8 text-sm" style={{ color: "var(--text-2)" }}>
          No account needed. Open the playground and start immediately.
        </p>
        <Link
          href="/playground"
          className="btn-primary inline-flex items-center gap-2 rounded-xl px-8 py-3 text-sm shadow-lg"
          style={{ boxShadow: "0 4px 24px rgba(236,72,153,0.25)" }}
        >
          <Play className="h-4 w-4" />
          Launch Playground
        </Link>
      </section>

    </div>
  );
}
