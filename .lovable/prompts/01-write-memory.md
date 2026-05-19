# Write Memory

**Triggers:** "write memory", "end memory", end-of-session checkpoint.

**Purpose:** Persist everything learned, done, and pending so the next AI session has zero context loss.

## Workflow (summary)

1. **Audit** — list what was done, what's pending, what was learned, what went wrong.
2. **Update `.lovable/memory/`** — append (never truncate) topic files. Update `memory/index.md` for any new file.
3. **Update `.lovable/plan.md`** — task statuses; move fully-complete items to `## Completed` in same file.
4. **Update `.lovable/suggestions.md`** — single file; move implemented items to `## Implemented Suggestions`.
5. **Update issues** — `pending-issues/XX-name.md` while open; move to `solved-issues/XX-name.md` with `## Solution`, `## Iteration Count`, `## Learning`, `## What NOT to Repeat` once resolved.
6. **Update `.lovable/strictly-avoid.md`** — for any pattern that must never be repeated.
7. **Update `.lovable/cicd-index.md` + `.lovable/cicd-issues/XX-issue-name.md`** — for any CI/CD pipeline issue (no duplicates).
8. **Validate** — index integrity, no orphans, no item in both pending and solved.
9. **Confirm** — emit the summary block from the canonical spec.

## Rules

- All `.md` files lowercase + hyphenated.
- Plans and suggestions are single files; do NOT fragment.
- Never overwrite blindly — read before write.
- Never delete history — mark done, move to `## Completed`.
- Never create `.lovable/memories/` (with trailing `s`). Correct path is `.lovable/memory/`.
- Never test `internal/...` packages (project-wide hard rule).
- Never run stateful git commands (sandbox-forbidden).

Full canonical spec is the user's "Write memory" prompt text.
