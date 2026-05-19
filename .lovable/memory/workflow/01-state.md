# Workflow State

Last updated: end of session 2026-05-19.

## Phase status

| # | Phase | Status |
|---|---|---|
| 0 | Stabilize build + 6 failing tests | ✅ Done (Tasks H–L) |
| 1 | `docs/ARCHITECTURE.md` | ✅ Done |
| 2 | `docs/LLM_GUIDELINE.md` | ✅ Done |
| 3 | Unit tests for 11 public packages | ✅ Done |
| 4 | `docs/extensibility.md` | ✅ Done |
| 5 | Generics refactor — strategy (c) freeze | ✅ Done |
| 6 | CMD extraction (`errcmdportable`) | ✅ Done; package-move + alias policy still ⏳ pending user input |
| 7 | Fix bad git remote | 🚫 Blocked — user-side only (`git remote set-url`) |
| M | Remove dead `core-v9` / `enum-v10` deps | ✅ Done |
| F | Re-run `.\run.ps1 -tc`, confirm 11/11 ✓ | 🔄 In Progress — last run 10/11 ⚠ REVIEW |

## Active blocker

Waiting for user to paste first 80 lines of `data/coverage/build-errors.txt` (or
fallback: `go build ./tests/errtypetests/... ./tests/errorwrappertests/... 2>&1 | Select -First 60`).

Without that log, the 57 cascade-blocked sub-packages cannot be diagnosed.
Known upstream drift candidates (see `blockers/02-upstream-api-drift.md`):
`corestr.New.LinkedCollections`, `converters.StringToIntegerWithDefault`,
`coredynamic.SliceItemsAsStringsAny`, `errwrappers.NewEmpty`,
`errtype.InvalidValidate`, `errnew.Type.Message`, `errnew.NotFound.Simple`,
plus a real `sync.noCopy` violation in `errwrappers/Collection.go:1216`.

## Next step for the next AI

1. If user pasted the build-errors log → diagnose root cause, patch, re-run `.\run.ps1 -tc`.
2. If user said `scan` → grep codebase for the drift signatures above and propose speculative patches.
3. If user said `done` → close roadmap, mark Phase F ✅ with caveat note.
4. Otherwise → restate blocker; do NOT loop with empty `next` responses.
