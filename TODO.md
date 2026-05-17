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

### Phase 5 — Generics refactor of `errdata/*`
- Introduce `errdata.Result[T any]` consolidating
  `errstr`, `errbool`, `errint`, `errint64`, `errfloat`, `errbytes`,
  `errbytesarr`, `errjsonresult`, `errslice`, etc.
- Decision pending: keep old packages as type aliases for one minor
  cycle, or remove outright. **Blocker**: user input.
- Migrate `LinkedCollections` constructor work into the new generic shape.

### Phase 6 — This file
- ✅ Created.
- Going forward: every PR that adds a `// TODO` in Go source must add a
  matching bullet here in the same commit.

### Phase 7 — Repository hygiene
- Fix the misconfigured git remote.
- **Blocker**: confirm working branch name + correct remote URL.

### Test coverage backlog
- Add the 12 new `tests/integratedtests/*tests` packages introduced in
  Phase 3 to `.\run.ps1 -tc` (Windows) and `./run.sh -tc` (Linux)
  coverage lists.
- Decision pending: incremental update vs. one batch commit.

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
