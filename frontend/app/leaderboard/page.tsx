// ─────────────────────────────────────────────
//  Leaderboard Page
// ─────────────────────────────────────────────

"use client";

import { useState, useEffect } from "react";
import { Trophy, Medal, Award, Loader2 } from "lucide-react";
import apiClient from "@/lib/api";

interface LeaderboardEntry {
  rank: number;
  user_id: string;
  username: string;
  problems_solved: number;
  total_submissions: number;
  acceptance_rate: number;
  score: number;
}

// Mock data for when the API isn't available yet
const MOCK_LEADERBOARD: LeaderboardEntry[] = [
  { rank: 1, user_id: "1", username: "alice", problems_solved: 3, total_submissions: 5, acceptance_rate: 80, score: 300 },
  { rank: 2, user_id: "2", username: "bob", problems_solved: 2, total_submissions: 8, acceptance_rate: 62.5, score: 200 },
  { rank: 3, user_id: "3", username: "charlie", problems_solved: 1, total_submissions: 3, acceptance_rate: 66.7, score: 100 },
];

function RankIcon({ rank }: { rank: number }) {
  if (rank === 1) return <Trophy className="h-5 w-5 text-yellow-400" />;
  if (rank === 2) return <Medal className="h-5 w-5 text-slate-300" />;
  if (rank === 3) return <Award className="h-5 w-5 text-amber-600" />;
  return <span className="text-slate-400 font-mono text-sm">{rank}</span>;
}

export default function LeaderboardPage() {
  const [entries, setEntries] = useState<LeaderboardEntry[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchLeaderboard() {
      try {
        const response = await apiClient.get("/leaderboard");
        setEntries(response.data);
      } catch {
        // Fall back to mock data if endpoint not available
        setEntries(MOCK_LEADERBOARD);
      } finally {
        setLoading(false);
      }
    }
    fetchLeaderboard();
  }, []);

  return (
    <div className="mx-auto max-w-4xl px-4 py-10">
      {/* Header */}
      <div className="mb-8 text-center">
        <div className="mb-4 inline-flex rounded-full bg-yellow-900/20 p-3">
          <Trophy className="h-8 w-8 text-yellow-400" />
        </div>
        <h1 className="text-3xl font-bold text-white">Leaderboard</h1>
        <p className="mt-2 text-slate-400">Top performers ranked by problems solved</p>
      </div>

      {/* Top 3 podium */}
      {!loading && entries.length >= 3 && (
        <div className="mb-10 flex items-end justify-center gap-4">
          {/* 2nd place */}
          <div className="flex flex-col items-center">
            <div className="mb-2 flex h-12 w-12 items-center justify-center rounded-full bg-slate-700 text-lg font-bold text-slate-200">
              {entries[1]?.username.charAt(0).toUpperCase()}
            </div>
            <p className="text-sm font-medium text-slate-300">{entries[1]?.username}</p>
            <p className="text-xs text-slate-500">{entries[1]?.problems_solved} solved</p>
            <div className="mt-2 flex h-16 w-20 items-center justify-center rounded-t-lg bg-slate-700 text-2xl font-bold text-slate-300">
              2
            </div>
          </div>

          {/* 1st place */}
          <div className="flex flex-col items-center">
            <Trophy className="mb-1 h-6 w-6 text-yellow-400" />
            <div className="mb-2 flex h-14 w-14 items-center justify-center rounded-full bg-yellow-900/40 border-2 border-yellow-500 text-xl font-bold text-yellow-300">
              {entries[0]?.username.charAt(0).toUpperCase()}
            </div>
            <p className="text-sm font-semibold text-white">{entries[0]?.username}</p>
            <p className="text-xs text-slate-400">{entries[0]?.problems_solved} solved</p>
            <div className="mt-2 flex h-24 w-20 items-center justify-center rounded-t-lg bg-yellow-900/30 border border-yellow-800/50 text-2xl font-bold text-yellow-400">
              1
            </div>
          </div>

          {/* 3rd place */}
          <div className="flex flex-col items-center">
            <div className="mb-2 flex h-12 w-12 items-center justify-center rounded-full bg-slate-700 text-lg font-bold text-amber-600">
              {entries[2]?.username.charAt(0).toUpperCase()}
            </div>
            <p className="text-sm font-medium text-slate-300">{entries[2]?.username}</p>
            <p className="text-xs text-slate-500">{entries[2]?.problems_solved} solved</p>
            <div className="mt-2 flex h-12 w-20 items-center justify-center rounded-t-lg bg-amber-900/20 text-2xl font-bold text-amber-600">
              3
            </div>
          </div>
        </div>
      )}

      {/* Full table */}
      <div className="rounded-xl border border-slate-800 bg-slate-900 overflow-hidden">
        <table className="w-full">
          <thead>
            <tr className="border-b border-slate-800 text-left">
              <th className="px-4 py-3 text-xs font-medium uppercase tracking-wider text-slate-500">Rank</th>
              <th className="px-4 py-3 text-xs font-medium uppercase tracking-wider text-slate-500">User</th>
              <th className="px-4 py-3 text-right text-xs font-medium uppercase tracking-wider text-slate-500">Solved</th>
              <th className="px-4 py-3 text-right text-xs font-medium uppercase tracking-wider text-slate-500">Submissions</th>
              <th className="px-4 py-3 text-right text-xs font-medium uppercase tracking-wider text-slate-500">Acceptance</th>
              <th className="px-4 py-3 text-right text-xs font-medium uppercase tracking-wider text-slate-500">Score</th>
            </tr>
          </thead>
          <tbody>
            {loading ? (
              <tr>
                <td colSpan={6} className="py-12 text-center">
                  <Loader2 className="mx-auto h-6 w-6 animate-spin text-slate-500" />
                </td>
              </tr>
            ) : entries.length === 0 ? (
              <tr>
                <td colSpan={6} className="py-12 text-center text-slate-500">
                  No submissions yet. Be the first!
                </td>
              </tr>
            ) : (
              entries.map((entry) => (
                <tr
                  key={entry.user_id}
                  className="border-b border-slate-800/50 transition-colors hover:bg-slate-800/30"
                >
                  <td className="px-4 py-3">
                    <div className="flex items-center justify-center w-6">
                      <RankIcon rank={entry.rank} />
                    </div>
                  </td>
                  <td className="px-4 py-3">
                    <div className="flex items-center gap-3">
                      <div className="flex h-8 w-8 items-center justify-center rounded-full bg-slate-700 text-sm font-medium text-slate-200">
                        {entry.username.charAt(0).toUpperCase()}
                      </div>
                      <span className="font-medium text-slate-200">{entry.username}</span>
                    </div>
                  </td>
                  <td className="px-4 py-3 text-right font-mono text-sm text-green-400">
                    {entry.problems_solved}
                  </td>
                  <td className="px-4 py-3 text-right font-mono text-sm text-slate-400">
                    {entry.total_submissions}
                  </td>
                  <td className="px-4 py-3 text-right font-mono text-sm text-slate-400">
                    {entry.acceptance_rate.toFixed(1)}%
                  </td>
                  <td className="px-4 py-3 text-right font-mono text-sm font-semibold text-white">
                    {entry.score}
                  </td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>
    </div>
  );
}
