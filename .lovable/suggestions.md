# Suggestions

## Active Suggestions

### Proactive drift scan
- **Status:** Pending (offered to user, awaiting `scan` keyword)
- **Priority:** High
- **Description:** Grep codebase for known drift signatures (`corestr.New.LinkedCollections`, `converters.StringToIntegerWithDefault`, `coredynamic.SliceItemsAsStringsAny`, `errwrappers.NewEmpty`, `errtype.InvalidValidate`, `errnew.Type.Message`, `errnew.NotFound.Simple`) and propose speculative patches without waiting for the build-errors log.
- **Added:** 2026-05-19

### Fix `sync.noCopy` violation in `errwrappers/Collection.go:1216`
- **Status:** Pending
- **Priority:** Medium
- **Description:** `go vet` flags a real noCopy violation. Independent of upstream drift; can be patched once the failure log confirms the exact call site.
- **Added:** 2026-05-19

### Decide CMD package move path + alias policy (Phase 6 leftover)
- **Status:** Pending (needs user decision)
- **Priority:** Low
- **Description:** `errcmd*` packages may eventually move to a clearer namespace; need user to ratify path + back-compat alias rules.
- **Added:** 2026-05-18

## Implemented Suggestions

### Remove dead `core-v9` / `enum-v10` dependencies (Task M)
- **Implemented:** 2026-05-18
- **Notes:** Verified no `.go` file imported either; removed from `go.mod` + `go.sum`. Eliminated a class of phantom compile errors.

### Freeze legacy `errdata/*`, adopt `erranygen` in new code (Phase 5)
- **Implemented:** 2026-05-18
- **Notes:** Banner added to 8 `Result.go` files; migration recipe in `docs/extensibility.md` §6.3.

### Split `errcmdportable.Detect()` by build tag (Tasks G + K)
- **Implemented:** 2026-05-18
- **Notes:** `detect_default.go` (js/wasip1) + `detect_os.go` (native, with inlined `autoOsRunner` to avoid import cycle).

### Add `errcmdbridge` + `reflectinternal` test coverage (Tasks I + J)
- **Implemented:** 2026-05-18
- **Notes:** Brings test-package count to 29. `reflectinternal` test lives under `tests/integratedtests/`, NOT inside `internal/`.
