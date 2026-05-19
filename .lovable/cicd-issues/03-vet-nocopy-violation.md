# 03 — `go vet` `sync.noCopy` violation

**Phase:** 7 (Vet) of `run.ps1 -tc`.

**Location:** `errwrappers/Collection.go:1216`.

**Cause:** Likely a struct embedding `sync.Mutex` (or RWMutex / WaitGroup) being passed or returned by value.

**Fix sketch:**
- Switch receivers to pointer receivers around that line.
- Or hold the mutex via pointer (`*sync.Mutex`) inside the struct.

**Status:** ⏳ Pending direct inspection.
