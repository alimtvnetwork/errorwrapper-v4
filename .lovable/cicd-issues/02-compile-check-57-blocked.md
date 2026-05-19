# 02 — Compile Check: 57 sub-packages blocked

**Phase:** 6 (Compile Check) of `run.ps1 -tc`.

**Symptom:** 57 cascade-blocked packages in the compile step. Suspected upstream API drift.

**Suspect signatures** (see `.lovable/memory/blockers/02-upstream-api-drift.md`):
`corestr.New.LinkedCollections`, `converters.StringToIntegerWithDefault`, `coredynamic.SliceItemsAsStringsAny`, `errwrappers.NewEmpty`, `errtype.InvalidValidate`, `errnew.Type.Message`, `errnew.NotFound.Simple`.

**Unblock:**
```powershell
Get-Content .\data\coverage\build-errors.txt | Select-Object -First 80
```

**Status:** ⏳ Awaiting build-errors log.
