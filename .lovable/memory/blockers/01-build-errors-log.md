# Blocker: Need `data/coverage/build-errors.txt`

Status: 🚫 Blocked on user input.

The last `.\run.ps1 -tc` run reported 57 cascade-blocked sub-packages in the
Compile Check phase. Without the captured stderr, root cause cannot be confirmed.

## What to ask the user

```powershell
Get-Content .\data\coverage\build-errors.txt | Select-Object -First 80
```

Fallback if file is empty:

```powershell
go build ./tests/errtypetests/... ./tests/errorwrappertests/... 2>&1 | Select-Object -First 60
```

## Alternative path

If the user says `scan`, proactively grep the codebase for the drift
signatures in `blockers/02-upstream-api-drift.md` and propose speculative patches.

## Do NOT

- Loop on bare `next` with the same message. State the blocker once, offer `scan` / `done` alternatives.
