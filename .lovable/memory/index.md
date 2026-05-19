# Memory Index

Institutional knowledge for errorwrapper-v3. Every file under `.lovable/memory/` must be listed here.

## Files

- [overview](./01-project-overview.md) — What errorwrapper-v3 is, layout, tooling.
- [workflow/state](./workflow/01-state.md) — Current phase status (Phases 0–7, Tasks H–M).
- [workflow/test-runner](./workflow/02-test-runner.md) — `run.ps1 -tc` pipeline, phase panel meaning.
- [decisions/01-phase5-freeze](./decisions/01-phase5-freeze.md) — Strategy (c): freeze legacy `errdata/*`, prefer `erranygen.Result[T]`.
- [decisions/02-no-internal-tests](./decisions/02-no-internal-tests.md) — Hard rule: never test `internal/...` packages.
- [decisions/03-edge-runtime-split](./decisions/03-edge-runtime-split.md) — `errcmdportable` build-tag split for js/wasip1.
- [conventions/01-go-layout](./conventions/01-go-layout.md) — Public vs internal packages, test location rules.
- [blockers/01-build-errors-log](./blockers/01-build-errors-log.md) — Waiting on first 80 lines of `data/coverage/build-errors.txt`.
- [blockers/02-upstream-api-drift](./blockers/02-upstream-api-drift.md) — Known drift signatures pending upstream coordination.
- [avoid/01-stateful-git](./avoid/01-stateful-git.md) — Agent cannot run git remote/push/pull mutations.
- [avoid/02-internal-package-tests](./avoid/02-internal-package-tests.md) — Forbidden by user; mirror of decision 02.

## Core rules (always applied)

- Never write tests for any `internal/...` package. Tests only live under `tests/integratedtests/<pkg>tests/` for public packages.
- Legacy `errdata/*` is frozen; new generic code uses `errdata/erranygen.Result[T]`.
- Agent cannot mutate git state (no `git remote`, `push`, `pull`, `commit`). Phase 7 is user-side.
- Always list remaining tasks at end of every response.
