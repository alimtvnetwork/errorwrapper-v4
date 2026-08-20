# Project Overview — errorwrapper-v4

Go library providing typed error wrappers, collections, command runners, and reflection helpers. Also ships a small TanStack Start landing site under `src/routes/`.

## Layout

- `errnew/`, `errtype/`, `errconv/`, `errwrappers/`, `eithererr/`, `trydo/`, `refs/`, `errverify/`, `errdefer/`, `errfunc/`, `errcmd/`, `linuxservicecmd/` — public packages.
- `errdata/{errany,errbool,errbyte,errfloat,errfloat64,errint,errjson,errstr,errcasted,erranygen}/` — typed result/collection variants. Legacy ones are **frozen**; `erranygen` is the generic successor.
- `errcmdportable/` — edge-safe Runner; `osadapter/` for native, `errcmdbridge/` for `*errcmd.Result` conversion.
- `internal/{consts,messages,reflectinternal}/` — implementation details. **Never tested directly.**
- `tests/integratedtests/<pkg>tests/` — all test files live here.
- `scripts/` + `run.ps1` / `run.sh` — local CI orchestrator with coverage.
- `src/` — TanStack Start landing site (routes: `/`, `/docs`, `/docs/extending-error-types`, `/docs/llm-guideline`).

## Tooling

- Go modules: `go.mod`. `core-v9` and `enum-v10` were removed (Task M) — dead deps.
- Local CI: `.\run.ps1 -tc` runs lint → tidy → build → tests → coverage in 11 phases.
- Frontend: bun, Vite 7, TanStack Start v1, Tailwind v4 (src/styles.css tokens).
