# TODO

Extraction marker for outstanding work. Generated from inline `TODO` /
`FIXME` markers in source plus open phase items from the architecture
roadmap. Keep this file in sync when you add or resolve a TODO comment.

> Convention: every inline `// TODO` in Go source should have a matching
> bullet here. When you remove the comment, remove the bullet.

---

## Inline TODOs in source

### `errdata/errstr/LinkedCollections.go:9`
```go
// LinkedCollections TODO constructors
type LinkedCollections struct { ... }
```
- **What**: Constructors are missing. Today callers must zero-build the
  struct and assign fields manually.
- **Action**: Add `NewLinkedCollections(...)`,
  `NewLinkedCollectionsWithError(...)` mirroring the patterns in sibling
  `errdata/errstr` files.
- **Phase**: Folds into Phase 5 (generics refactor) — the new
  `errdata.Result[T]` should subsume this.

### `errnew/constructors.go:122` and `:133` — `NotImpl` / `NotImplPtrUsingStackSkip`
```go
"TODO: url(" + url + ")"
```
- **What**: The error message hard-codes the literal string `TODO: url(...)`.
  These are intentional placeholders surfaced to callers of `NotImpl`.
- **Action**: Treat as **WONTFIX** — the literal `TODO:` prefix is the
  signal that this code path is a stub. Do not remove without updating
  every consumer that pattern-matches on the prefix.

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
- `LinkedCollections` constructor work folds into whichever path wins.

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

### Worker / edge-runtime compatibility (deferred research)
- `errcmd` spawns real OS processes and is incompatible with
  Cloudflare Workers / serverless runtimes. Investigate a portable
  façade that no-ops or returns a typed `NotSupported` wrapper on
  non-OS targets. Tracked in `docs/extensibility.md` §6.

### Streaming verifier (deferred research)
- `errverify.CollectionVerifier` materializes the full slice. For very
  large collections a streaming variant would be useful. Tracked in
  `docs/extensibility.md` §6.

---

## Done (historical, kept for grep-ability)

- ✅ Phase 0 — Stabilize build
- ✅ Phase 1 — Architecture map
- ✅ Phase 2 — LLM guideline
- ✅ Phase 3 — Core unit tests (11 packages + `errverify` review +
  `errcmd` pure-utility tests)
- ✅ Phase 4 — `docs/extensibility.md`
- ✅ Phase 5 PoC — `errdata/erranygen.Result[T]` + Convey tests
- ✅ Phase 6 — `TODO.md` extraction marker (this file)
- ✅ Coverage backlog — auto-discovered, no manual list to maintain
