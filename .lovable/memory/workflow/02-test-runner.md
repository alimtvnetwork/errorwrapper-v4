# Test Runner — `run.ps1 -tc`

11-phase pipeline (PowerShell + Go). Last observed result: **10/11 ✓, STATUS ⚠ REVIEW**.

## Phases (in order)

1. Git pull (currently warns: bad remote → 404, see Phase 7).
2. Dependency fetch (`go mod download`).
3. Clean `data/` folder.
4. In-package import lint.
5. `go mod tidy` check.
6. Compile check (target: 27/27 or 29/29 packages).
7. Vet.
8. Unit tests (`go test ./tests/integratedtests/...`).
9. Coverage run (2 packages).
10. Coverage report generation.
11. Final summary panel.

## Useful artifacts

- `data/coverage/build-errors.txt` — captured stderr from compile phase.
- `data/test-logs/failing-tests.txt` — verbose failing test output.
- `data/coverage/*.out` — Go coverage profiles.

## Unblock commands (for the user)

```powershell
Get-Content .\data\coverage\build-errors.txt | Select-Object -First 80
# fallback:
go build ./tests/errtypetests/... ./tests/errorwrappertests/... 2>&1 | Select-Object -First 60
```
