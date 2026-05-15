"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { Terminal, BookOpen, Github, Menu, X } from "lucide-react";
import { useState } from "react";
import clsx from "clsx";

const NAV_LINKS = [
  { href: "/playground", label: "Playground", icon: Terminal },
  { href: "/problems",   label: "Problems",   icon: BookOpen  },
];

export default function Navbar() {
  const pathname = usePathname();
  const [open, setOpen] = useState(false);

  return (
    <header
      className="sticky top-0 z-50 backdrop-blur-md border-b"
      style={{ background: "rgba(13,10,14,0.88)", borderColor: "var(--border)" }}
    >
      <div className="mx-auto flex h-14 max-w-7xl items-center justify-between px-5">

        {/* ── Logo ─────────────────────────── */}
        <Link href="/" className="flex items-center gap-2.5">
          <div
            className="flex h-7 w-7 items-center justify-center rounded-lg"
            style={{
              background: "linear-gradient(135deg, rgba(236,72,153,0.2), rgba(245,158,11,0.2))",
              border: "1px solid rgba(236,72,153,0.3)",
            }}
          >
            <Terminal className="h-3.5 w-3.5" style={{ color: "#ec4899" }} />
          </div>
          <span className="text-base font-bold tracking-tight gradient-text">Executo</span>
        </Link>

        {/* ── Desktop nav ──────────────────── */}
        <nav className="hidden items-center gap-0.5 md:flex">
          {NAV_LINKS.map(({ href, label, icon: Icon }) => {
            const active = pathname === href || pathname.startsWith(href + "/");
            return (
              <Link
                key={href}
                href={href}
                className="relative flex items-center gap-1.5 rounded-lg px-3.5 py-2 text-sm font-medium transition-colors"
                style={{ color: active ? "var(--text)" : "var(--text-2)" }}
              >
                <Icon className="h-3.5 w-3.5" />
                {label}
                {/* Active underline — gradient line, not background fill */}
                {active && (
                  <span
                    className="absolute bottom-0 left-3 right-3 gradient-line"
                  />
                )}
              </Link>
            );
          })}
        </nav>

        {/* ── Right ────────────────────────── */}
        <div className="flex items-center gap-3">
          <a
            href="https://github.com"
            target="_blank"
            rel="noopener noreferrer"
            className="btn-ghost hidden items-center gap-1.5 rounded-lg px-3 py-1.5 text-xs md:flex"
          >
            <Github className="h-3.5 w-3.5" />
            GitHub
          </a>

          {/* Online indicator */}
          <div className="hidden items-center gap-1.5 text-xs md:flex" style={{ color: "var(--text-3)" }}>
            <span
              className="h-1.5 w-1.5 rounded-full animate-pulse-dot"
              style={{ background: "var(--success)" }}
            />
            Online
          </div>

          {/* Mobile toggle */}
          <button
            onClick={() => setOpen(!open)}
            className="rounded-lg p-1.5 transition-colors md:hidden"
            style={{ color: "var(--text-2)" }}
          >
            {open ? <X className="h-5 w-5" /> : <Menu className="h-5 w-5" />}
          </button>
        </div>
      </div>

      {/* ── Mobile menu ──────────────────────── */}
      {open && (
        <div
          className="border-t px-4 py-3 md:hidden"
          style={{ borderColor: "var(--border)", background: "var(--surface)" }}
        >
          <nav className="flex flex-col gap-1">
            {NAV_LINKS.map(({ href, label, icon: Icon }) => {
              const active = pathname === href || pathname.startsWith(href + "/");
              return (
                <Link
                  key={href}
                  href={href}
                  onClick={() => setOpen(false)}
                  className={clsx(
                    "flex items-center gap-2 rounded-lg px-3 py-2.5 text-sm font-medium transition-colors"
                  )}
                  style={{ color: active ? "var(--text)" : "var(--text-2)" }}
                >
                  <Icon className="h-4 w-4" />
                  {label}
                </Link>
              );
            })}
          </nav>
        </div>
      )}
    </header>
  );
}
