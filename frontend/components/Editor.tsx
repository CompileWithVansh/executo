// ─────────────────────────────────────────────
//  Monaco Editor Wrapper Component
// ─────────────────────────────────────────────

"use client";

import { useRef, useCallback } from "react";
import MonacoEditor, { OnMount, BeforeMount } from "@monaco-editor/react";
import type * as Monaco from "monaco-editor";

interface EditorProps {
  /** Monaco language identifier (e.g. "python", "java", "cpp") */
  language: string;
  /** Current editor value */
  value: string;
  /** Called when the editor content changes */
  onChange?: (value: string | undefined) => void;
  /** Whether the editor is read-only */
  readOnly?: boolean;
  /** Editor height (default: 100%) */
  height?: string;
  /** Minimum number of lines to show */
  minLines?: number;
}

// ── Editor Theme ──────────────────────────────
// Custom dark theme matching Executo's design system

const EXECUTO_THEME: Monaco.editor.IStandaloneThemeData = {
  base: "vs-dark",
  inherit: true,
  rules: [
    // Rose Gold Dark syntax — restrained, readable
    { token: "comment",           foreground: "6b7280", fontStyle: "italic" },
    { token: "keyword",           foreground: "ec4899" },          // rose
    { token: "keyword.control",   foreground: "ec4899" },
    { token: "storage",           foreground: "ec4899" },
    { token: "string",            foreground: "fca5a5" },          // muted peach
    { token: "string.escape",     foreground: "f59e0b" },
    { token: "number",            foreground: "f59e0b" },          // gold
    { token: "constant.numeric",  foreground: "f59e0b" },
    { token: "type",              foreground: "f9a8d4" },          // soft rose
    { token: "class",             foreground: "f9a8d4" },
    { token: "function",          foreground: "fde68a" },          // warm gold
    { token: "entity.name",       foreground: "fde68a" },
    { token: "variable",          foreground: "f5f5f5" },          // near-white
    { token: "variable.parameter",foreground: "e4e4e7" },
    { token: "operator",          foreground: "a1a1aa" },
    { token: "punctuation",       foreground: "71717a" },
    { token: "tag",               foreground: "ec4899" },
    { token: "attribute.name",    foreground: "f59e0b" },
    { token: "attribute.value",   foreground: "fca5a5" },
  ],
  colors: {
    // Backgrounds — layered depth
    "editor.background":                     "#0d0a0e",
    "editor.foreground":                     "#f5f5f5",
    "editor.lineHighlightBackground":        "#1a121d",
    "editor.lineHighlightBorder":            "#00000000",

    // Selection — subtle rose tint
    "editor.selectionBackground":            "#ec489928",
    "editor.inactiveSelectionBackground":    "#ec489914",
    "editor.selectionHighlightBackground":   "#ec489910",

    // Line numbers
    "editorLineNumber.foreground":           "#3d2d42",
    "editorLineNumber.activeForeground":     "#6b7280",

    // Cursor — rose
    "editorCursor.foreground":               "#ec4899",
    "editorCursor.background":               "#0d0a0e",

    // Find
    "editor.findMatchBackground":            "#f59e0b30",
    "editor.findMatchHighlightBackground":   "#f59e0b18",
    "editor.findMatchBorder":                "#f59e0b60",

    // Widgets (autocomplete, hover)
    "editorWidget.background":               "#1a121d",
    "editorWidget.border":                   "#2a1f2d",
    "editorWidget.foreground":               "#f5f5f5",
    "editorSuggestWidget.background":        "#1a121d",
    "editorSuggestWidget.border":            "#2a1f2d",
    "editorSuggestWidget.foreground":        "#f5f5f5",
    "editorSuggestWidget.selectedBackground":"#ec489918",
    "editorSuggestWidget.highlightForeground":"#ec4899",

    // Scrollbar
    "scrollbarSlider.background":            "#2a1f2d80",
    "scrollbarSlider.hoverBackground":       "#3d2d42aa",
    "scrollbarSlider.activeBackground":      "#3d2d42cc",

    // Indent guides
    "editorIndentGuide.background":          "#2a1f2d",
    "editorIndentGuide.activeBackground":    "#3d2d42",

    // Bracket match
    "editorBracketMatch.background":         "#ec489918",
    "editorBracketMatch.border":             "#ec489960",

    // Minimap
    "minimap.background":                    "#0d0a0e",
    "minimap.selectionHighlight":            "#ec489930",
  },
};

// ── Default Editor Options ────────────────────

const DEFAULT_OPTIONS: Monaco.editor.IStandaloneEditorConstructionOptions = {
  fontSize: 14,
  fontFamily: "'JetBrains Mono', 'Fira Code', 'Cascadia Code', Consolas, monospace",
  fontLigatures: true,
  lineHeight: 22,
  tabSize: 4,
  insertSpaces: true,
  wordWrap: "on",
  minimap: { enabled: false },
  scrollBeyondLastLine: false,
  renderLineHighlight: "line",
  cursorBlinking: "smooth",
  cursorSmoothCaretAnimation: "on",
  smoothScrolling: true,
  formatOnPaste: true,
  formatOnType: false,
  autoIndent: "full",
  bracketPairColorization: { enabled: true },
  guides: {
    bracketPairs: true,
    indentation: true,
  },
  suggest: {
    showKeywords: true,
    showSnippets: true,
  },
  quickSuggestions: {
    other: true,
    comments: false,
    strings: false,
  },
  padding: { top: 16, bottom: 16 },
  scrollbar: {
    verticalScrollbarSize: 6,
    horizontalScrollbarSize: 6,
  },
  overviewRulerLanes: 0,
  hideCursorInOverviewRuler: true,
  renderWhitespace: "none",
  contextmenu: true,
  mouseWheelZoom: true,
};

// ── Component ─────────────────────────────────

export default function Editor({
  language,
  value,
  onChange,
  readOnly = false,
  height = "100%",
}: EditorProps) {
  const editorRef = useRef<Monaco.editor.IStandaloneCodeEditor | null>(null);

  // Register custom theme before editor mounts
  const handleBeforeMount: BeforeMount = useCallback((monaco) => {
    monaco.editor.defineTheme("executo-dark", EXECUTO_THEME);

    // Add Python-specific completions
    monaco.languages.registerCompletionItemProvider("python", {
      provideCompletionItems: (model, position) => {
        const word = model.getWordUntilPosition(position);
        const range = {
          startLineNumber: position.lineNumber,
          endLineNumber: position.lineNumber,
          startColumn: word.startColumn,
          endColumn: word.endColumn,
        };
        return {
          suggestions: [
            {
              label: "List[int]",
              kind: monaco.languages.CompletionItemKind.TypeParameter,
              insertText: "List[int]",
              range,
            },
            {
              label: "Dict[str, int]",
              kind: monaco.languages.CompletionItemKind.TypeParameter,
              insertText: "Dict[str, int]",
              range,
            },
          ],
        };
      },
    });
  }, []);

  const handleMount: OnMount = useCallback((editor) => {
    editorRef.current = editor;

    // Focus the editor on mount
    editor.focus();

    // Add keyboard shortcut: Ctrl+Enter to submit (fires a custom event)
    editor.addCommand(
      // Monaco.KeyMod.CtrlCmd | Monaco.KeyCode.Enter
      2048 | 3,
      () => {
        const event = new CustomEvent("editor:submit");
        window.dispatchEvent(event);
      }
    );
  }, []);

  return (
    <div className="h-full w-full bg-slate-950">
      <MonacoEditor
        height={height}
        language={language}
        value={value}
        theme="executo-dark"
        options={{
          ...DEFAULT_OPTIONS,
          readOnly,
        }}
        onChange={onChange}
        beforeMount={handleBeforeMount}
        onMount={handleMount}
        loading={
          <div className="flex h-full items-center justify-center bg-slate-950">
            <div className="text-sm text-slate-500">Loading editor...</div>
          </div>
        }
      />
    </div>
  );
}
