# Decision: Minor-version bump on every code change

**Date:** 2026-05-19
**Source:** User directive.

## Rule

Any change to project code (Go packages, frontend TS/TSX, scripts, configs that ship) MUST bump at least the **minor** version in every place a version string lives.

## Exception

Files under `.release/` are out of bounds — never read, modify, or version them.

## Application checklist

- `package.json` — `version` field
- `go.mod` — module comment (if version embedded)
- `version.go` / `consts.go` constants exposing a version
- `CHANGELOG.md` — new section header

## Why

Forces visible movement on every diff so downstream consumers (and other AIs reviewing history) can never miss that "something changed."
