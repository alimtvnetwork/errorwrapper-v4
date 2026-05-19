# Decision: `errcmdportable` Build-Tag Split

Date: 2026-05-18 (Tasks G + K).

## Problem

`errcmdportable.Detect()` needs to return an `os/exec`-backed Runner on native OS, but edge targets (`js`, `wasip1`) must not pull `os/exec`.

## Solution

Two files in `errcmdportable/`:

- `detect_default.go` — build tag `//go:build js || wasip1`. Returns `NoProcessRunner`.
- `detect_os.go` — build tag `//go:build !js && !wasip1`. Inlines a tiny `autoOsRunner` (NOT delegating to `osadapter.New()` — that created an import cycle, see Task K).

`errcmdportable/osadapter/` remains available for explicit opt-in callers.

## What NOT to do

- Do not have `detect_os.go` import `osadapter` — `osadapter` imports `errcmdportable` for the `Runner` interface → cycle.
- Do not put the OS branch in `Runner.go` without a build tag — edge bundlers will pull `os/exec`.
