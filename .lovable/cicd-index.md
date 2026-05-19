# CI/CD Issues Index

Tracking issues discovered in the local CI pipeline (`run.ps1` / `run.sh`).

| # | Issue | Status |
|---|---|---|
| [01](./cicd-issues/01-git-pull-404.md) | `git pull` phase 404s on missing remote | 🚫 Blocked — user-side |
| [02](./cicd-issues/02-compile-check-57-blocked.md) | Compile Check: 57 sub-packages blocked | ⏳ Pending build-errors.txt |
| [03](./cicd-issues/03-vet-nocopy-violation.md) | `go vet` flags `sync.noCopy` at `errwrappers/Collection.go:1216` | ⏳ Pending fix |
| [04](./cicd-issues/04-phase-status-review.md) | Final summary: 10/11 phases ✓, STATUS ⚠ REVIEW | 🔄 Resolves when 02 + 03 fixed |
