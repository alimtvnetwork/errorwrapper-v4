# Pending Issues — Single Tracker

One row per unresolved issue. Update in-place. Move resolved entries to the `## Resolved` section at the bottom; never delete.

---

## Open

### 01 · build-errors-log-missing
- severity: blocker
- description: `data/coverage/build-errors.txt` not pasted; cannot diagnose 57 cascade-blocked sub-packages.
- owner: user
- unblockCommand: `Get-Content .\data\coverage\build-errors.txt | Select-Object -First 80`
- fallback: `go build ./tests/errtypetests/... ./tests/errorwrappertests/... 2>&1 | Select-Object -First 60`

### 02 · git-remote-404
- severity: low
- description: `run.ps1` Phase 1 fails: remote `https://github.com/alimtvnetwork/errorwrapper-v4.git/` returns 404.
- owner: user (sandbox forbids stateful git)
- fix: `git remote set-url origin <correct-url> && git fetch origin`

### 03 · sync-nocopy-violation
- severity: medium
- description: `go vet` flags `sync.noCopy` at `errwrappers/Collection.go:1216`.
- owner: agent
- fix: switch to pointer receiver / pointer-to-mutex field.

### 04 · specs-missing
- severity: blocker (planning)
- description: No `spec/` folder exists; user assumes specs are present. Risk report cannot evaluate spec quality.
- owner: user
- fix: clarify location (cross-project? external?) or authorize authoring from scratch.

### 05 · go-toolchain-unavailable-in-sandbox
- severity: medium (workflow)
- description: `go` binary not installed in Lovable sandbox; agent cannot compile, vet, or test Go code locally. Only the TanStack frontend builds.
- owner: agent-environment
- fix: user runs Go commands locally and pastes output, OR install go via nix (`nix run nixpkgs#go -- build ./...`) per-call.

---

## Resolved

(none yet — Phase 0 fixes pre-date this tracker; see `workflow/01-state.md`.)
