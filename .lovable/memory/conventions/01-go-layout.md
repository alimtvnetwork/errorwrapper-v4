# Go Layout Conventions

## Public packages

- Live at repo root or one level deep (e.g. `errnew/`, `errdata/errstr/`).
- Tests for them live in `tests/integratedtests/<pkg>tests/`, NEVER alongside the source.

## Internal packages

- `internal/...` — no tests, no exported surface assumed.
- Exercised indirectly via public-package tests.

## Test packages

- Convention: `tests/integratedtests/<pkg>tests/<Thing>_test.go`.
- Use `goconvey` — top-level `Convey(...)` calls REQUIRE `t` as the second argument.
- Use `ShouldNotPointTo` for pointer-identity checks, not `ShouldNotEqual` (which DeepEquals).
- `EmptyPtr()` returns `nil` by design → assert `ShouldBeNil`.

## Frozen markers

Legacy `errdata/*/Result.go` files carry `// Frozen: prefer erranygen.Result[T] for new code`.
Do not extend frozen types; add new functionality to `erranygen`.
