# Project Plan — errorwrapper-v3 + Lovable Frontend

Authoritative roadmap for AI handoff. Mirrors `.lovable/plan.md` (legacy in-folder copy) but lives at repo root for discoverability.

> **Versioning rule:** Any code change (Go or frontend) MUST bump at least the **minor** version everywhere a version string appears — except files under `.release/` which are out of bounds.

---

## Phase Map

| Phase | Title | Status |
|---|---|---|
| 0 | Stabilize build + 6 failing tests | ✅ Done |
| 1 | `docs/ARCHITECTURE.md` | ✅ Done |
| 2 | `docs/LLM_GUIDELINE.md` | ✅ Done |
| 3 | Unit tests for 11 public packages | ✅ Done (29 test-pkgs) |
| 4 | `docs/extensibility.md` | ✅ Done |
| 5 | Generics refactor (strategy (c) freeze) | ✅ Done |
| 6 | CMD extraction (`errcmdportable`) | ✅ Done; alias policy pending |
| 7 | Fix bad git remote | 🚫 User-side |
| M | Remove dead `core-v9` / `enum-v10` deps | ✅ Done |
| F | Re-run `run.ps1 -tc` → 11/11 ✓ | 🔄 Blocked on 02/03 |
| S | Author or import `spec/` folder | 🚫 Blocked on user |

---

## Prioritized Backlog

### P0 — Unblock implementation
1. **Clarify spec location** (`specs-missing`)
   - Objective: Confirm whether specs exist in another project or need to be authored here.
   - Dependencies: user input.
   - Expected outputs: `spec/` directory populated, OR a written authoring plan.
   - Acceptance: README index + data-model + API surface + acceptance criteria present.

2. **Paste / capture build-errors log** (`build-errors-log-missing`)
   - Objective: Diagnose 57 blocked sub-packages.
   - Dependencies: user runs PowerShell command and pastes top 80 lines.
   - Outputs: confirmed list of broken API call sites.
   - Acceptance: agent produces a patch plan covering ≥80% of blocked packages.

### P1 — Hands-on fixes
3. **Fix `sync.noCopy` at `errwrappers/Collection.go:1216`**
   - Objective: Pass `go vet ./...` cleanly.
   - Dependencies: none (independent).
   - Outputs: patch to `Collection.go` + minor version bump.
   - Acceptance: `go vet ./...` clean; `go build ./errwrappers/...` clean.

4. **Patch upstream API drift** (after P0 #2)
   - Objective: Restore 27/27 or 29/29 compile success.
   - Dependencies: build-errors log.
   - Outputs: per-package patches + minor version bumps.
   - Acceptance: `.\run.ps1 -tc` reaches Phase 6 with 0 failures.

### P2 — Housekeeping
5. **`errcmd*` package move + alias policy** (`cmd-package-move-policy`)
   - Objective: Finalize naming.
   - Dependencies: user decision on target namespace.
   - Outputs: file moves + type-alias compat layer + minor version bump.
   - Acceptance: old import paths still compile via aliases.

6. **Git remote fix** (user-only)
   - Acceptance: `git pull` Phase 1 of `run.ps1 -tc` succeeds.

### P3 — Frontend (TanStack Start)
7. **Decide frontend purpose**
   - Options: (a) docs site for the Go library, (b) interactive demo/playground, (c) unrelated app, (d) delete frontend scaffold.
   - Dependencies: user decision.
   - Outputs: replace `data-lovable-blank-page-placeholder` in `src/routes/index.tsx`.
   - Acceptance: real homepage shipped; SEO `<head>` populated.

---

## Next Task Selection

Pick ONE of the following to start implementation (ordered by readiness):

- **A** — Fix `sync.noCopy` in `errwrappers/Collection.go:1216` (only requires Go toolchain; can do via `nix run nixpkgs#go`).
- **B** — Author `spec/` folder skeleton from scratch (no external blockers; needs scope from user).
- **C** — Build frontend homepage (no Go dependency; pure TanStack work).
- **D** — Paste build-errors.txt → start patching drift (highest value, requires user paste).

---

## Versioning Checklist (run on every code change)

- [ ] Bump minor version in `package.json`
- [ ] Bump minor version in `go.mod` module comment (if present)
- [ ] Bump minor version in any `version.go` / `consts.go` constants
- [ ] Bump minor version in `CHANGELOG.md` header
- [ ] **Do NOT** touch anything under `.release/`

---

## Handoff Contract (for the next AI)

1. Read `.lovable/memory/index.md` first.
2. Read `.lovable/memory/workflow/01-state.md` for phase status.
3. Read `.lovable/memory/suggestions/01-suggestions.md` for active suggestions.
4. Read `.lovable/memory/pending-issues/01-pending-issues.md` for blockers.
5. Read this `plan.md` for backlog priority.
6. Respect `.lovable/strictly-avoid.md` — non-negotiable.
7. Ask the user which Next Task Selection item (A/B/C/D) to start.
