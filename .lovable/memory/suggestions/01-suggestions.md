# Suggestions — Active Tracker

Single-file aggregate. One row per suggestion. Update status in-place; do not split into multiple files.

Statuses: `open` · `inProgress` · `done` · `dropped`

File naming convention for any future split (only if this file exceeds ~500 lines):
`YYYYMMDD-HHMMSS-suggestion-<slug>.md`

---

## Schema (apply to every entry)

```
### <slug>
- suggestionId: <slug>
- createdAt: YYYY-MM-DD
- source: Lovable
- affectedProject: errorwrapper-v4 | frontend | both
- description: <one line>
- rationale: <why it matters>
- proposedChange: <what to do, where>
- acceptanceCriteria: <how we know it's done>
- status: open | inProgress | done | dropped
- completionNotes: <filled when done>
```

---

## Active

### proactive-drift-scan
- suggestionId: proactive-drift-scan
- createdAt: 2026-05-19
- source: Lovable
- affectedProject: errorwrapper-v4
- description: Grep codebase for known upstream drift signatures and propose speculative patches without waiting for build-errors.txt.
- rationale: 57 sub-packages are cascade-blocked; user hasn't supplied the build log for 5+ turns. Speculative grep can unblock progress.
- proposedChange: Run `rg` for `corestr.New.LinkedCollections`, `converters.StringToIntegerWithDefault`, `coredynamic.SliceItemsAsStringsAny`, `errwrappers.NewEmpty`, `errtype.InvalidValidate`, `errnew.Type.Message`, `errnew.NotFound.Simple`; map call sites; produce a patch plan.
- acceptanceCriteria: Patch plan covers ≥80% of the 57 blocked packages OR confirms the drift signatures are absent.
- status: open
- completionNotes:

### fix-sync-nocopy-collection
- suggestionId: fix-sync-nocopy-collection
- createdAt: 2026-05-19
- source: Lovable
- affectedProject: errorwrapper-v4
- description: `go vet` flags `sync.noCopy` at `errwrappers/Collection.go:1216`.
- rationale: Real bug, independent of upstream drift, blocks vet phase.
- proposedChange: Inspect line 1216; switch to pointer receiver or pointer-to-mutex; re-run `go vet ./errwrappers/...`.
- acceptanceCriteria: `go vet ./...` returns 0 noCopy diagnostics.
- status: open
- completionNotes:

### cmd-package-move-policy
- suggestionId: cmd-package-move-policy
- createdAt: 2026-05-18
- source: User (Phase 6 leftover)
- affectedProject: errorwrapper-v4
- description: Decide final namespace for `errcmd*` packages + back-compat alias policy.
- rationale: Phase 6 functionally complete; only naming/alias decision pending.
- proposedChange: User ratifies path; agent then performs move + adds type aliases at old locations.
- acceptanceCriteria: Move executed, alias file created, old import paths still compile.
- status: open
- completionNotes:

### write-or-import-specs
- suggestionId: write-or-import-specs
- createdAt: 2026-05-19
- source: Lovable
- affectedProject: both
- description: There is no `spec/` folder in the repo. User assumes one exists. Either import the missing specs from another project or author them here.
- rationale: Without specs, any AI handoff has near-zero chance of building "the right thing"; risk report cannot evaluate spec quality.
- proposedChange: User clarifies whether specs live in another @-mentioned project (cross_project tools available) or need to be written from scratch.
- acceptanceCriteria: `spec/` directory exists at repo root with at least: README index, data model, API surface, acceptance criteria.
- status: open
- completionNotes:

---

## Implemented (do not delete; for traceability)

### remove-dead-deps
- status: done · 2026-05-18 · `core-v9` + `enum-v10` removed from `go.mod`/`go.sum`.

### freeze-errdata-legacy
- status: done · 2026-05-18 · Banner added to 8 `Result.go` files; recipe in `docs/extensibility.md` §6.3.

### errcmdportable-build-tag-split
- status: done · 2026-05-18 · `detect_default.go` (js/wasip1) + `detect_os.go` (native).

### errcmdbridge-and-reflectinternal-coverage
- status: done · 2026-05-18 · 29 test-packages total; `reflectinternal` test lives under `tests/integratedtests/` (NOT inside `internal/`).
