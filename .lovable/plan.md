# Completion Plan — errorwrapper-v3

Three roadmap items remain: **Phase 0** (3 runtime test failures), **Phase 5** (generics refactor of `errdata/*`), **Phase 7** (bad git remote). This plan splits the remaining work into 6 sequenced execution phases (A–F) so each is independently shippable and verifiable.

---

## Phase A — Diagnose the 3 runtime failures

**Goal:** Get exact failure messages for `Test_EmptyPtr_HasError`, `Test_ClonePtr`, `TestErrconv_GetPtr`.

Steps:
1. User pastes `data/test-logs/failing-tests.txt` (or runs `go test -run "Test_EmptyPtr_HasError|Test_ClonePtr|TestErrconv_GetPtr" ./... -v` and pastes output).
2. Confirm whether failures stem from:
   - assertion semantics (e.g. `ShouldNotEqual` doing DeepEqual on `*Wrapper`),
   - actual logic regression in `EmptyPtr` / `ClonePtr` / `errconv.GetPtr`,
   - or a transitive break from the recent `errcmdtests` import fix.

**Exit:** Root cause identified for all 3, written into `TODO.md` Phase 0 section.

---

## Phase B — Fix the 3 runtime failures

**Goal:** Green build with 0 runtime test failures and 0 blocked packages.

Steps:
1. Apply targeted fixes per Phase A diagnosis:
   - If assertion semantics → switch to `ShouldNotPointTo` / address comparison.
   - If logic regression → patch `Wrapper.go` / `errconv/Get.go` and add regression tests.
2. Re-run `.\run.ps1 -tc`.
3. Verify: 27/27 compile, 0 runtime failures, `errconvtests` + `errorwrappertests` green.

**Exit:** All Phase 0 items checked off in `TODO.md`.

---

## Phase C — Phase 5 strategy decision (1 question)

**Goal:** Pick the migration strategy for `errdata/*` → `erranygen.Result[T]`.

Ask the user once:
- **(a) Type-alias bridge** — `type Result = erranygen.Result[string]`. Smallest diff, but breaks method sets that depend on concrete receivers.
- **(b) Embed** — each `errstr.Result` embeds `erranygen.Result[string]`. Preserves identity, allows extra methods, modest churn.
- **(c) Freeze legacy + generics-only in new code** *(recommended)* — zero churn to legacy callers, new packages adopt `erranygen` directly.

**Exit:** Strategy locked. `docs/extensibility.md` §6.3 records the decision + rationale.

---

## Phase D — Execute Phase 5 migration

Branches by strategy chosen in Phase C.

### If (c) freeze — smallest path
1. Add `// Frozen: prefer erranygen.Result[T] for new code` banner to `errdata/*/Result.go` files.
2. Add `docs/extensibility.md` §6.3 with the migration recipe + worked example.
3. Done.

### If (b) embed — medium path
1. For each of `errstr`, `errint`, `errbool`, `errbyte`, `errfloat`, `errfloat64`, `errany`:
   - Add embedded `*erranygen.Result[T]` field on `Result`.
   - Forward `Value`/`ErrorWrapper` access through the embed.
   - Update constructors in `New.Result.*` to populate both.
2. Run unit tests after each package; commit per package.
3. Update `docs/extensibility.md` §6.3.

### If (a) alias — largest blast radius
1. Replace concrete `Result` structs with `type Result = erranygen.Result[T]`.
2. Move package-specific methods to free functions or wrapper types.
3. Fix all downstream call sites (expect 50–200 edits).
4. Full test sweep.

**Exit:** All Phase 5 items checked off in `TODO.md`; new code path documented.

---

## Phase E — Phase 7 git remote fix (user action)

**Goal:** `git pull` in `run.ps1` succeeds.

Steps (user-side, agent cannot run git state mutations):
1. User confirms the correct remote URL (likely a typo in the org/repo name on GitHub).
2. User runs locally:
   ```
   git remote set-url origin <correct-url>
   git fetch origin
   ```
3. Re-run `.\run.ps1 -tc` and confirm "Pulling latest from remote" succeeds.

**Exit:** Phase 7 checked off; `TestRunnerCore.psm1 → Invoke-GitPull` no longer warns.

---

## Phase F — Final verification & sign-off

1. Run `.\run.ps1 -tc` end-to-end.
2. Confirm in the phase summary panel:
   - Compile Check: 27/27
   - Runtime failures: 0
   - All previously-blocked subpackages now contribute to coverage
3. Update `TODO.md`: mark Phases 0, 5, 7 ✅; close the roadmap.
4. Write a short `CHANGELOG.md` entry summarizing the stabilization + generics adoption.

**Exit:** Roadmap complete.

---

## Full roadmap status after this plan

| # | Phase | Status |
|---|---|---|
| 0 | Stabilize build + 6 failing tests | A + B |
| 1 | `docs/ARCHITECTURE.md` | ✅ done |
| 2 | `docs/LLM_GUIDELINE.md` | ✅ done |
| 3 | Unit tests for 11 public packages | ✅ done (compile clears after B) |
| 4 | `docs/extensibility.md` | ✅ done |
| 5 | Generics refactor of `errdata/*` | C + D |
| 6 | CMD extraction (errcmdportable) | ✅ done |
| 7 | Fix bad git remote | E |
| — | Final verification | F |

---

## Execution order

A → B → C → D → E (parallel, user-side) → F

Phases A, C, E need user input; B, D, F are agent-executed.

**On `next`** without further input, agent will request the failing-tests log (Phase A) and propose strategy **(c)** for Phase C.
