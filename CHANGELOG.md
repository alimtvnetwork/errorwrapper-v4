# Changelog

## 2026-05-18 — Task L: runtime test failures fixed

### Fixed
- `errconv/GetPtr` — non-`*errorwrapper.Wrapper` pointers (notably
  `*errwrappers.Collection`) were being short-circuited to `EmptyPtr()`
  inside the `IsPointer` branch, so they never reached the
  `BasicErrWrapper` / `BaseErrorOrCollectionWrapper` switch. Now only
  typed-nil `*Wrapper` returns early; other pointer types fall through.
  Fixes `TestErrconv_GetPtr` line 115. Also clears `TestErrconv_Get`,
  which was failing via shared panic from the reflect test below.
- `tests/integratedtests/reflectinternaltests/reflect_test.go:28` —
  `GetElementTypeMaxTry(&[]int{1}, 0)` peels one pointer layer then
  exhausts budget and returns `nil`. Test now asserts `ShouldBeNil`
  instead of dereferencing `.Kind()` on nil.

## 2026-05-18 — Roadmap stabilization & generics adoption

### Fixed (Phase 0)
- `tests/integratedtests/errcmdtests/Utilities_test.go` — corrected bad import
  path (`.../errorwrapper-v4/errorwrapper` → `.../errorwrapper-v4`).
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

### Added (Task J — 2026-05-18)
- `tests/integratedtests/reflectinternaltests/reflect_test.go` — first test
  coverage for `internal/reflectinternal` (~450 lines of reflection + `unsafe`
  pointer code). Convey suite covers `GetElementType{,MaxTry}`,
  `GetElementTypesMaxTry`, `GetTypeName`, `IsType`, `IsTypeSame`,
  `GetPointerInfo`, `IsBytesOrBytesPointer`, `IsStringOrStringPointer`,
  `IsString`, `IsStringsOrStringsPointer`, `IsIntegersOrIntegersPointer`,
  `IsIntegerOrIntegerPointer`, `IsInteger`, `IsBoolean`, `IsBooleanPointer`,
  `IsFloat64sOrFloat64sPointer`, `NewScanReport`, and `GetFieldValue` —
  value, pointer, and non-matching paths each asserted. Brings
  test-package count to 29.

### Fixed (Task K — 2026-05-18, import cycle)
- `errcmdportable/detect_os.go` — Task G introduced an import cycle:
  the file lived in package `errcmdportable` and imported the
  `errcmdportable/osadapter` subpackage, which itself imports
  `errcmdportable` for the `Runner` interface + `Result` type.
  Fixed by inlining a tiny `autoOsRunner` (build-tag-guarded by
  `//go:build !js && !wasip1`) directly in the file instead of
  delegating to `osadapter.New()`. Edge builds still drop `os/exec`
  via the build tag; `osadapter` remains available for explicit
  opt-in callers. Verified with `go build ./...`.

### Pending (user action)
- **Phase 7** — fix bad git remote: `git remote set-url origin <correct-url>`
  (current origin 404s on `github.com/alimtvnetwork/errorwrapper-v4`).
- **Phase F** — re-run `.\run.ps1 -tc`. The Task K fix unblocks
  `errcmdportabletests` + `errcmdbridgetests`. Remaining blocked packages
  are pre-existing upstream API drift against `core-v9 v1.5.8`
  (`corestr.New.LinkedCollections`, `converters.StringToIntegerWithDefault`,
  `coredynamic.SliceItemsAsStringsAny`, `errwrappers.NewEmpty`,
  `errtype.InvalidValidate`, `errnew.Type.Message`, `errnew.NotFound.Simple`)
  plus a real `sync.noCopy` violation in `errwrappers/Collection.go:1216`
  — outside agent scope, needs upstream library coordination.



