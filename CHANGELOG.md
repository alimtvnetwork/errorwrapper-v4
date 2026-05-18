# Changelog

## 2026-05-18 — Roadmap stabilization & generics adoption

### Fixed (Phase 0)
- `tests/integratedtests/errcmdtests/Utilities_test.go` — corrected bad import
  path (`.../errorwrapper-v3/errorwrapper` → `.../errorwrapper-v3`).
  Unblocked `go mod tidy` and 49 cascade-blocked packages.
- `Test_EmptyPtr_HasError` — `EmptyPtr()` returns `nil` by design;
  assertion switched to `ShouldBeNil`.
- `Test_ClonePtr` — replaced `ShouldNotEqual` (DeepEqual) with
  `ShouldNotPointTo` (pointer identity).
- `TestErrconv_GetPtr` + `TestErrconv_Get` — added missing `t` argument
  to 10 top-level `Convey(...)` calls.

### Changed (Phase 5 — strategy (c) freeze)
- `errdata/{errany,errbool,errbyte,errfloat,errfloat64,errint,errjson,errstr}/Result.go`
  — added `// Frozen: prefer erranygen.Result[T] for new code` banner.
- Legacy `errdata/*` packages remain fully supported; new code should adopt
  `errdata/erranygen.Result[T]` directly.

### Documented
- `docs/extensibility.md` §6.3 — Phase 5 decision record + migration recipe.
- `TODO.md` — roadmap closeout entry.

### Fixed (Task H — 2026-05-18)
- `errnew/constructors.go` — replaced placeholder `"TODO: url("+url+")"` messages
  in `NotImpl()` and `NotImplPtrUsingStackSkip()` with proper `"Not implemented: "+url`.

### Changed (Task G — 2026-05-18)
- `errcmdportable/Detect()` split into build-tag-guarded files:
  `detect_default.go` (js/wasip1 → `NoProcessRunner`) and `detect_os.go`
  (native OS → `osadapter.New()`). Edge builds stay `os/exec`-free;
  native builds auto-wire the production adapter. Removed the TODO from
  `Runner.go` package doc.

### Added (Task I — 2026-05-18)
- `tests/integratedtests/errcmdportabletests/errcmdbridgetests/Bridge_test.go`
  — first test coverage for `errcmdportable/errcmdbridge.FromErrcmdResult`:
  nil input → zero `Result`, successful `*errcmd.Result` → stdout pass-through,
  and `errorWrapper` carry-over as `Result.Err` with stderr trimming. Brings
  test-package count to 28.

### Pending (user action)
- **Phase 7** — fix bad git remote: `git remote set-url origin <correct-url>`
  (current origin 404s on `github.com/alimtvnetwork/errorwrapper-v3`).
- **Phase F** — re-run `.\run.ps1 -tc` and confirm 28/28 compile,
  0 runtime failures.
- **Task J** — add tests for `internal/reflectinternal` (~450 lines of
  reflection + `unsafe` pointer code, zero coverage today).

