# `sync.noCopy` violation in `errwrappers/Collection.go:1216`

## Description
`go vet` flags a real `sync.noCopy` violation at `errwrappers/Collection.go:1216`. Independent of the upstream drift.

## Root Cause
Under investigation — likely a struct embedding `sync.Mutex` (or similar) being passed/returned by value somewhere around line 1216.

## Steps to Reproduce
1. `go vet ./errwrappers/...`
2. Observe noCopy diagnostic at line 1216

## Attempted Solutions
- [ ] Not yet attempted; needs the build-errors log or direct inspection.

## Priority
Medium — does not block compile, does fail vet phase.

## Blocked By
Nothing strictly; can be tackled as soon as the user runs the next session.
