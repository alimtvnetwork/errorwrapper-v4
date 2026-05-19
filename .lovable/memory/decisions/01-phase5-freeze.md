# Decision: Phase 5 — Strategy (c) "Freeze + Generics-Only in New Code"

Date: 2026-05-18.

## Context

Three options considered for migrating legacy `errdata/{errstr,errbool,errint,...}` to `erranygen.Result[T]`:

- (a) Type alias — strips type-specific methods (`IsTrue`, `Int`, `Bool`, `SplitLines`, ...). Breaking.
- (b) Embed generic in legacy struct — changes JSON shape, breaks serialized data.
- (c) Freeze legacy, adopt `erranygen` in new code only — zero churn. **Chosen.**

## Implementation

- Banner added to `errdata/{errany,errbool,errbyte,errfloat,errfloat64,errint,errjson,errstr}/Result.go`:
  `// Frozen: prefer erranygen.Result[T] for new code`
- Migration recipe in `docs/extensibility.md` §6.3.

## Future direction

Greenfield work uses `erranygen.Result[T]` directly. No legacy callers need to migrate.
