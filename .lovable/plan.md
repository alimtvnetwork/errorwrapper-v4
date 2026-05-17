# ErrorWrapper-v3 Work Plan

Based on your two summaries, here is a step-by-step plan. I will execute **one phase per `next`**. If something is ambiguous in a phase, I will ask before doing it.

---

## Phase 0 — Stabilize current build (prerequisite)
The 6 failing tests from the last `.\run.ps1 -tc` run still need confirmation on your Mac. Before adding new work we should:
- Confirm `git status` / latest commit is in sync with the fixes already applied (errverify IDs, label brackets, macOS skip).
- Re-run `.\run.ps1 -tc` and capture `data/test-logs/failing-tests.txt` if anything still fails.

## Phase 1 — Codebase read-through + architecture map
- Map the public packages: `errorwrapper` (root), `errnew`, `errtype`, `errwrappers`, `errverify`, `errcmd`, `errdata/*`, `errdefer`, `errfunc(s)`, `trydo`, `eithererr`, `linuxservicecmd`, `ref(s)`, `errwonce`.
- Mark `internal/*` as **no-test** (per your rule).
- Produce a short `docs/ARCHITECTURE.md` describing Wrapper, Collection, References, StackSkip, ErrorType enum, Compile pipeline, CMD subsystem.

## Phase 2 — LLM guideline document
- Create `docs/LLM_GUIDELINE.md` covering:
  - Naming (`errnew.X` vs `errorwrapper.X`), why no `errwrapper.ErrorWrapper`.
  - When to use `errnew.Type`, `errnew.Messages`, `errnew.NotFound`, etc.
  - StackSkip rules (utility funcs use `codestack.Skip1`).
  - Collection vs single Wrapper.
  - Result[T] usage pattern.
  - Forbidden patterns / circular-dep avoidance via `errnew`.

## Phase 3 — Unit tests for public packages
Iterate package-by-package, skipping `internal/*`. Order:
1. `errtype` (enum, variation, mapping)
2. `errnew` (all `new*Creator` constructors)
3. `errorwrapper` root (Wrapper, ConcatNew, Compile, References, StackTrace)
4. `errwrappers` (Collection, MutexCollection, StateCounter)
5. `refs` / `ref`
6. `errdefer`, `errfunc`, `errfuncs`
7. `trydo`, `eithererr`
8. `errdata/*` (after Phase 5 generics refactor — see note)
9. `errverify`
10. `errcmd` + `linuxservicecmd` (platform-guarded)
11. `errwonce`

## Phase 4 — Extensibility research: user-injectable error types
Deliver `docs/EXTENDING_ERROR_TYPES.md` with **2–3 concrete approaches**, pros/cons, sample code:
- **A. Registry pattern** — `errtype.Register(name, code, msg)` returning a `Variation`.
- **B. Interface-based custom types** — user implements `errtype.CustomVariation` interface.
- **C. Code-gen** — extend the `all-generate.go` generator to merge user YAML/JSON enum file at build time.

We pick one together before implementing.

## Phase 5 — Generics refactor of `errdata/*`
Collapse `errany`, `errbool`, `errbyte`, `errfloat`, `errfloat64`, `errint`, `errstr`, `errjson` Result/Results into a single generic package:
- `errdata/errgeneric/Result[T]`, `Results[T]`, `Result2[T,U]`, `ResultWithApplicable[T]`, `ResultsWithErrorCollection[T]`.
- Keep old packages as thin type aliases for backward compatibility (e.g. `type Result = errgeneric.Result[string]`).
- Update creators accordingly.

## Phase 6 — CMD package extraction (TODO marker only)
- Add `TODO.md` note to move `errcmd/` and `linuxservicecmd/` to a standalone module in the future.
- No code move now.

## Phase 7 — Branch hygiene
- Fix the bad git remote (currently points to `alimtvnetwork/errorwrapper-v3` which 404s) so `git pull` stops failing in `run.ps1`. Confirm the correct remote URL with you first.

---

## Open questions (please answer before Phase 1 if possible)
1. **Branch name** — what should I name the working branch?
2. **Correct git remote URL** for the repo (current one 404s).
3. **Phase 5 backward compatibility** — OK to keep old `errdata/errstr` etc. as aliases, or remove entirely?
4. **Test runner** — should new tests be added to `.\run.ps1 -tc` coverage list, or kept separate until ready?

Reply `next` to start **Phase 0**, or answer the open questions first.