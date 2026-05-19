# 01 — `git pull` phase 404s

**Phase:** 1 (Git pull) of `run.ps1 -tc`.

**Symptom:**
```
remote: Repository not found.
fatal: repository 'https://github.com/alimtvnetwork/errorwrapper-v3.git/' not found
✗ git pull failed (continuing anyway) (source: TestRunnerCore.psm1 → Invoke-GitPull)
```

**Cause:** Wrong remote URL.

**Fix (user-side only — agent cannot run git state mutations):**
```sh
git remote set-url origin <correct-url>
git fetch origin
```

**Status:** 🚫 Blocked on user.
