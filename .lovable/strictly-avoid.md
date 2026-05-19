# Strictly Avoid

Hard prohibitions for this project. Never violate.

- **Tests inside `internal/...`** — Forbidden by user; coverage must come via public packages. See: `.lovable/memory/avoid/02-internal-package-tests.md`
- **Stateful git commands** (`push`, `pull`, `remote set-url`, `commit`, ...) — Sandbox-disabled; Phase 7 is user-side only. See: `.lovable/memory/avoid/01-stateful-git.md`
- **Extending frozen `errdata/*` types** — Strategy (c) freeze; add new functionality to `errdata/erranygen` instead. See: `.lovable/memory/decisions/01-phase5-freeze.md`
- **`errcmdportable/detect_os.go` importing `osadapter`** — Creates import cycle. Inline the OS runner instead. See: `.lovable/memory/decisions/03-edge-runtime-split.md`
- **Looping on `next` with identical "I'm blocked" output** — State the blocker once, offer `scan` / `done` alternatives, then stop repeating.
