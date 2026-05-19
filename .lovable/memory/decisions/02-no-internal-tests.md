# Decision: Never Test `internal/...` Packages

User-mandated, stressed repeatedly. Hard rule.

## Rule

- Zero `_test.go` files inside `internal/consts/`, `internal/messages/`, `internal/reflectinternal/`, or any future `internal/...` package.
- If any are found, delete them.
- Coverage for internal logic happens **indirectly** via tests of the public packages that consume it.
  Example: `internal/reflectinternal` is exercised through `tests/integratedtests/reflectinternaltests/` ONLY because that test lives under the integrated test tree and imports the public façade — it does NOT live inside `internal/`.

## Why

Internal packages are implementation details. Pinning tests to them prevents
refactor freedom and leaks private API into the test surface.
