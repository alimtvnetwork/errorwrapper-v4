# TODO

Extraction marker for outstanding work. Generated from inline `TODO` /
`FIXME` markers in source plus open phase items from the architecture
roadmap. Keep this file in sync when you add or resolve a TODO comment.

> Convention: every inline `// TODO` in Go source should have a matching
> bullet here. When you remove the comment, remove the bullet.

---

## Inline TODOs in source

- None remaining.

---

## Roadmap items (carried from Phase plan)

### Phase 5 — Generics refactor of `errdata/*` (PoC delivered)
- ✅ `errdata/erranygen.Result[T any]` added as additive PoC with tests.
- ⚠️ Full migration of legacy `errdata/{errstr,errbool,errint,…}` is
  **blocked**. Three viable strategies, each requires user buy-in:
  1. **Type alias** (e.g. `type Result = erranygen.Result[bool]`) —
     STRIPS all type-specific methods (`IsTrue`, `Int`, `Bool`,
     `SplitLines`, `ValidValue`, `SimpleStringOnce`, …). Breaking.
  2. **Embed generic in legacy struct** — changes JSON shape (nested
     fields), silently breaks any persisted serialized data.
  3. **Greenfield rewrite under new package path** — essentially what
     the PoC already is; legacy packages stay until callers migrate.
- Recommended path forward: keep PoC as the migration target, mark
  legacy packages "frozen", let new code adopt `erranygen` directly.
- ✅ `LinkedCollections` constructors added (mirrors `LinkedList` pattern).

### Phase 6 — This file
- ✅ Created.
- Going forward: every PR that adds a `// TODO` in Go source must add a
  matching bullet here in the same commit.

### Phase 7 — Repository hygiene (OUT OF AGENT SCOPE)
- ❌ Cannot be performed from the Lovable sandbox. The current `origin`
  is a Lovable-managed JWT-auth GCP endpoint; stateful git commands
  (`remote set-url`, push, etc.) are disabled per platform rules.
- **Action for the human maintainer**: after exporting the repo
  locally, run `git remote set-url origin <upstream-url>` and push.

### Test coverage backlog
- ✅ No action needed. `scripts/CoverageRunner.psm1` auto-discovers
  test packages via `go list ./tests/integratedtests/...` — the 12
  Phase 3 packages are picked up automatically by `run.ps1 -tc` /
  `run.sh -tc`.

### Worker / edge-runtime compatibility (delivered)
- ✅ `errcmdportable/Runner.go` — portable `Runner`, `NoProcessRunner`,
  `ErrNotSupported`, GOOS-based `Detect()`.
- ✅ `errcmdportable/osadapter/Runner.go` — real `os/exec` adapter in a
  subpackage so edge bundlers don't transitively pull `os/exec`.
- ✅ `errcmdportable/errcmdbridge/Bridge.go` — `*errcmd.Result` →
  portable `Result` converter.
- ✅ `tests/integratedtests/errcmdportabletests/Runner_test.go` —
  Convey suite (NoProcess sentinel, Detect default, osadapter exec +
  non-zero exit, edge-target skip).
- ✅ `docs/extensibility.md` §6 written.

### Streaming verifier (delivered)
- ✅ `errverify/StreamingCollectionVerifier.go` — Feed/Finish pull API,
  five match modes (Equal, EqualFold, Contains, ContainsFold, Regex via
  cached `regexp.Compile`), `FromVariant` mapper for upstream
  `stringcompareas.Variant`, mismatch aggregation, missing-expected
  detection, optional `ExpectedLength` soft check.
- ✅ `tests/integratedtests/errverifytests/StreamingCollectionVerifier_test.go`
  — Convey suite covering all 5 modes, extra/missing lines, length
  check, nil-source guard, invalid-regex setup error.
- ⬜ Follow-up: wire into `errwrappers.Collection` as a streaming
  consumer adapter; doc in `docs/extensibility.md` §6.

---

## Done (historical, kept for grep-ability)

- ✅ Phase 0 — Stabilize build
- ✅ Phase 1 — Architecture map
- ✅ Phase 2 — LLM guideline
- ✅ Phase 3 — Core unit tests (11 packages + `errverify` review +
  `errcmd` pure-utility tests)
- ✅ Phase 4 — `docs/extensibility.md`
- ✅ Phase 5 PoC — `errdata/erranygen.Result[T]` + Convey tests
- ✅ Phase 5a — `errdata/errstr.LinkedCollections` constructors + tests
- ✅ Phase 5b — Unit tests for all legacy `errdata/*` packages:
  `errbool`, `errbyte`, `errint`, `errfloat`, `errfloat64`, `errany`,
  `errstr` (Result/Results/Result2/ResultWithApplicable/
  ResultWithApplicable2/ResultsWithErrorCollection/Collection/
  LinkedList/SimpleStringOnce/LinkedCollections + constructors),
  `errjson` (Result/ResultsCollection + constructors), `errcasted`.
- ✅ Phase 5c — Constructor tests for `errdata/*` New/Empty creators.
- ✅ Phase 6 — `TODO.md` extraction marker (this file)
- ✅ Coverage backlog — auto-discovered, no manual list to maintain
