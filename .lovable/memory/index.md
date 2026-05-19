# Memory Index

> Read this FIRST every session. Every file under `.lovable/memory/` MUST be listed here.

## Core (always applied)

- NEVER write or modify tests for any package under `internal/...`. Tests are ONLY for public/external packages.
- NEVER run stateful git commands (`push`, `pull`, `remote set-url`, `commit`, ...) — sandbox-forbidden.
- ANY code change bumps at least the **minor** version everywhere except `.release/` (which is out of bounds).
- Suggestions tracked in ONE file: `memory/suggestions/01-suggestions.md`.
- Pending issues tracked in ONE file: `memory/pending-issues/01-pending-issues.md`.
- File naming: lowercase + hyphen, `XX-` prefix where order matters.

## Files

### Project
- [01-project-overview.md](./01-project-overview.md) — what this repo is.

### Workflow
- [workflow/01-state.md](./workflow/01-state.md) — phase status, last-known build result.
- [workflow/02-test-runner.md](./workflow/02-test-runner.md) — `run.ps1 -tc` 11-phase pipeline.

### Decisions
- [decisions/01-phase5-freeze.md](./decisions/01-phase5-freeze.md) — freeze `errdata/*`; use `erranygen` for new code.
- [decisions/02-no-internal-tests.md](./decisions/02-no-internal-tests.md) — hard rule, restated.
- [decisions/03-edge-runtime-split.md](./decisions/03-edge-runtime-split.md) — `errcmdportable` build-tag split.
- [decisions/04-versioning-policy.md](./decisions/04-versioning-policy.md) — minor bump on every change; `.release/` off-limits.

### Conventions
- [conventions/01-go-layout.md](./conventions/01-go-layout.md) — Go package layout rules.

### Avoid
- [avoid/01-stateful-git.md](./avoid/01-stateful-git.md)
- [avoid/02-internal-package-tests.md](./avoid/02-internal-package-tests.md)

### Blockers
- [blockers/01-build-errors-log.md](./blockers/01-build-errors-log.md) — need build-errors.txt.
- [blockers/02-upstream-api-drift.md](./blockers/02-upstream-api-drift.md) — suspected API drift signatures.

### Trackers (single-file aggregates)
- [suggestions/01-suggestions.md](./suggestions/01-suggestions.md) — all active + implemented suggestions.
- [pending-issues/01-pending-issues.md](./pending-issues/01-pending-issues.md) — all open + resolved issues.

## Out-of-folder references
- Repo-root `plan.md` — authoritative roadmap.
- `.lovable/strictly-avoid.md` — hard prohibitions summary.
- `.lovable/overview.md` — folder layout doc.
