# Git remote 404 (Phase 7)

## Description
`run.ps1` git-pull phase fails: `remote: Repository not found. fatal: repository 'https://github.com/alimtvnetwork/errorwrapper-v3.git/' not found`.

## Root Cause
Wrong/typo remote URL configured locally.

## Steps to Reproduce
1. Run `.\run.ps1 -tc`
2. Phase 1 (git pull) emits `✗ git pull failed (continuing anyway)`

## Attempted Solutions
- (none possible from agent — stateful git mutations are sandbox-forbidden)

## Priority
Low — pipeline continues, all other phases run.

## Blocked By
User must run locally: `git remote set-url origin <correct-url>` then `git fetch origin`.
