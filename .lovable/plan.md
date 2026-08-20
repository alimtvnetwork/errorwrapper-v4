# Completion Plan — errorwrapper-v4

Active roadmap. Fully-complete items live in `## Completed` at the bottom — never deleted.

## Active

### Phase F — Final verification (🔄 In Progress)
- Re-run `.\run.ps1 -tc` after Phase 7 + CI issues 02 + 03 resolved.
- Target: 11/11 phases ✓, 27/27 (or 29/29) compile, 0 runtime failures.

### Phase 7 — Git remote fix (🚫 Blocked — user-side)
- User runs locally: `git remote set-url origin <correct-url>` then `git fetch origin`.
- Then re-run `.\run.ps1 -tc`.
- Agent cannot perform this (sandbox forbids stateful git).

### CI/CD #02 — Compile Check: 57 blocked sub-packages (⏳ Pending log)
- Needs first 80 lines of `data/coverage/build-errors.txt`.
- Fallback: `go build ./tests/errtypetests/... ./tests/errorwrappertests/... 2>&1 | Select-Object -First 60`.
- Or user says `scan` → agent does speculative drift-signature grep.

### CI/CD #03 — `sync.noCopy` at `errwrappers/Collection.go:1216` (⏳ Pending)
- Switch to pointer receivers or pointer-to-mutex field around that line.

### Task #6 — CMD package move path + alias policy (⏳ Awaiting user decision)
- Optional housekeeping.

---

## Completed

### Phase 0 — Stabilize build ✅ (2026-05-18, Tasks H–L)
- Fixed bad import in `tests/integratedtests/errcmdtests/Utilities_test.go`.
- `Test_EmptyPtr_HasError` → `ShouldBeNil`.
- `Test_ClonePtr` → `ShouldNotPointTo`.
- `TestErrconv_Get*` → added missing `t` arg to top-level `Convey(...)` calls.
- `errconv/GetPtr` — fixed early return for non-`*Wrapper` pointers.
- `reflectinternaltests/reflect_test.go:28` — `ShouldBeNil` instead of dereferencing nil.

### Phase 1 — `docs/ARCHITECTURE.md` ✅
### Phase 2 — `docs/LLM_GUIDELINE.md` ✅
### Phase 3 — Unit tests for 11 public packages ✅ (29 test-packages total after Tasks I+J)
### Phase 4 — `docs/extensibility.md` ✅

### Phase 5 — Generics refactor ✅ (strategy (c) freeze)
- Banner added to 8 `errdata/*/Result.go` files.
- `docs/extensibility.md` §6.3 records the decision + migration recipe.

### Phase 6 — CMD extraction ✅
- `errcmdportable/Runner.go` + `NoProcessRunner` + GOOS-based `Detect()`.
- `osadapter/Runner.go` real `os/exec` adapter.
- `errcmdbridge/Bridge.go` converter + test coverage.

### Task G — `errcmdportable` build-tag split ✅
- `detect_default.go` (js/wasip1) + `detect_os.go` (native).

### Task H — `errnew/constructors.go` ✅
- Replaced TODO placeholder messages in `NotImpl()` / `NotImplPtrUsingStackSkip()`.

### Task I — `errcmdbridge` test coverage ✅
### Task J — `internal/reflectinternal` indirect test coverage ✅
- Test lives under `tests/integratedtests/reflectinternaltests/`, NOT inside `internal/`.

### Task K — Fix import cycle from Task G ✅
- `detect_os.go` inlines `autoOsRunner` instead of delegating to `osadapter.New()`.

### Task L — Runtime test failures ✅
- `errconv/GetPtr` short-circuit bug + `reflect_test.go:28` nil deref.

### Task M — Remove dead `core-v9` / `enum-v10` deps ✅
- No `.go` file imported them; removed from `go.mod` + `go.sum`.
