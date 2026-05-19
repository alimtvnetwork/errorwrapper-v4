# Avoid: Stateful Git Commands

The Lovable sandbox forbids `git add`, `commit`, `checkout`, `merge`, `rebase`,
`reset`, `stash`, `push`, `pull`, `remote set-url`, etc.

## Implication

Phase 7 ("fix bad git remote") **cannot** be done by the agent. User must run
locally:

```sh
git remote set-url origin <correct-url>
git fetch origin
```

The current `origin` 404s on `https://github.com/alimtvnetwork/errorwrapper-v3.git`.

## Symptom to recognize

`run.ps1` prints `remote: Repository not found.` then `✗ git pull failed (continuing anyway)`.
That is expected until the user fixes the remote — do not try to "fix" it from the agent side.
