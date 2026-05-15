import type { Metadata } from "next";
import "./globals.css";
import Navbar from "@/components/Navbar";

export const metadata: Metadata = {
  title: {
    default: "Executo — Online Code Runner",
    template: "%s | Executo",
  },
  description:
    "Run code instantly in Python, Java, C++, and JavaScript. No setup, no installs.",
  keywords: ["online compiler", "code runner", "python", "java", "cpp", "javascript"],
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <head>
        <link rel="preconnect" href="https://fonts.googleapis.com" crossOrigin="anonymous" />
        <link
          href="https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500;600&display=swap"
          rel="stylesheet"
        />
      </head>
      <body className="min-h-screen antialiased">
        <Navbar />
        <main className="min-h-[calc(100vh-3.5rem)]">{children}</main>
        <footer
          className="border-t py-5 text-center text-xs"
          style={{ borderColor: "var(--border)", color: "var(--text-3)" }}
        >
          Executo · Online Code Runner · Powered by Judge0
        </footer>
      </body>
    </html>
  );
}
