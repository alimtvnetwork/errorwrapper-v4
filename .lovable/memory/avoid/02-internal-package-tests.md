# Avoid: Tests Inside `internal/...`

Mirror of `decisions/02-no-internal-tests.md`. User has stressed this repeatedly.

- Never create `_test.go` inside any `internal/...` package.
- If you find one, delete it.
- Test internals indirectly through public-package test suites.
