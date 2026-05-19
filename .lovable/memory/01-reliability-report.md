# Reliability & Failure-Chance Report — Spec Handoff to Another AI

**Date:** 2026-05-19
**Scope:** Evaluating readiness of this repository's specifications for handoff to an arbitrary AI implementer.
**Verdict:** 🔴 **NOT READY** — primary blocker is structural, not editorial.

---

## 0. Honest Finding (read this first)

There is **no `spec/` or `specs/` folder in this repository.** I searched: nothing.

What *does* exist is operational/process memory in `.lovable/`:
- Phase tracker (`workflow/01-state.md`)
- CI pipeline doc (`workflow/02-test-runner.md`)
- Architectural decisions (4 files in `decisions/`)
- Blockers and avoid-lists
- A roadmap (`plan.md`)

These describe **how the project is being built**, not **what the product is supposed to do**. They are insufficient as a spec set. Any AI handed this repo today would have to infer the product from Go source code — a high-risk path.

Two doc files referenced by memory (`docs/ARCHITECTURE.md`, `docs/LLM_GUIDELINE.md`, `docs/extensibility.md`) are also **not present** in the working tree — they are marked "✅ Done" in the phase tracker but the `docs/` folder doesn't exist. This is a memory↔reality drift bug.

---

## 1. Success Probability Estimates

| Tier | Description | Success chance with current artifacts | Assumptions |
|---|---|---|---|
| **Trivial** | Single-file mechanical fix (e.g. `sync.noCopy` at known line) | **80%** | AI has Go toolchain; reads the strictly-avoid list. |
| **Simple** | Single-package API drift patch (signature rename) | **55%** | AI has build-errors.txt; doesn't over-refactor. |
| **Medium** | Cross-package refactor (e.g. CMD move + aliases) | **25%** | AI infers intent from memory notes; high chance of inventing wrong namespace. |
| **Complex** | Implementing a new feature end-to-end | **5–10%** | No spec exists. AI would invent requirements. |
| **End-to-end** | Build the "frontend purpose" (docs vs demo vs unrelated) | **<5%** | Pure guesswork without user input. |

**Headline number:** for anything beyond trivial fixes, a fresh AI has a **≤25% chance of producing what you actually want** on the first pass.

---

## 2. Failure Map

| Where | Why | Symptom |
|---|---|---|
| `spec/` directory | Doesn't exist | AI fabricates requirements; output diverges from intent. |
| `docs/*.md` files | Listed as "Done" in memory but missing from disk | AI follows phantom references; trust in memory collapses. |
| Upstream API drift | 57 packages blocked; no build log | AI patches wrong signatures, breaks more code. |
| `errwrappers/Collection.go:1216` | `sync.noCopy` flagged but root cause not characterized | AI applies the wrong fix (e.g. removes the mutex). |
| Frontend `src/routes/index.tsx` | Still `data-lovable-blank-page-placeholder` | AI builds a generic landing page that doesn't match the project. |
| `.release/` folder | "Out of bounds" rule lives only in user's head until now | AI touches it. (Now documented in `decisions/04-versioning-policy.md`.) |
| Cross-project assumption | User believes specs exist "across all projects" | AI cannot access without cross-project @mentions; will hallucinate. |
| Versioning rule | Brand new; no enforcement | AI forgets to bump minor version. |
| `internal/` test rule | Strongly stated in memory; survives | Low risk if AI reads `index.md`. |

---

## 3. Corrective Actions (prioritized)

| # | Action | Where | Expected reliability gain |
|---|---|---|---|
| 1 | **Author or import `spec/`** — at minimum: README, data-model.md, api-surface.md, acceptance-criteria.md per public package. | repo root `spec/` | +40–60% on Medium/Complex tiers. |
| 2 | **Recover or regenerate `docs/`** — `ARCHITECTURE.md`, `LLM_GUIDELINE.md`, `extensibility.md`. Memory says they exist; they don't. | `docs/` | +15% (restores memory↔reality trust). |
| 3 | **Capture build-errors.txt** — paste top 80 lines OR run fallback `go build` command. | attach to chat | +25% on Simple/Medium drift fixes. |
| 4 | **Add a one-paragraph "product north star"** to `.lovable/memory/01-project-overview.md` answering: who uses errorwrapper-v3, what one problem does it solve, what's out of scope. | memory | +10% across all tiers. |
| 5 | **Frontend intent decision** — pick (a) docs site, (b) demo, (c) unrelated, (d) delete. | chat answer + `plan.md` update | +20% on frontend work. |
| 6 | **Versioning hook** — add a pre-commit check OR a `CHANGELOG.md` template line that fails CI if untouched. | repo root | +5% (catches forgotten bumps). |
| 7 | **`.release/` guard** — add to `.lovable/strictly-avoid.md` and `.gitignore`-style read protection note. | already done in `decisions/04` | +5%. |

---

## 4. Readiness Decision

**Ready for implementation?** No.

**Minimum bar to start handing tasks to another AI:**

1. ✅ Suggestion + pending-issue trackers exist (just done — single file each).
2. ✅ `plan.md` exists at repo root with Next Task Selection.
3. ✅ Versioning policy documented.
4. ❌ `spec/` folder with at least one project's data-model + acceptance criteria. **← blocker**
5. ❌ Build-errors.txt captured OR `scan` directive given. **← blocker**
6. ❌ Frontend purpose declared. **← blocker (frontend work only)**

For **trivial Go fixes** (e.g. `sync.noCopy`), the bar is already met — that work can start now.

---

## 5. Questions & Concerns (per user request)

1. **Where are the specs?** You referenced "spec folder content for all projects" — none exists in this repo. Are they in another Lovable project I should `@mention`, in an external system, or do they need to be authored from scratch here?
2. **What is errorwrapper-v3 actually for?** Memory describes phases and tests but no product north star. One-sentence answer would unblock 40% of risk.
3. **Frontend scope** — is the TanStack Start scaffold meant to become a docs site, a demo, or should it be deleted?
4. **`.release/` folder** — it doesn't currently exist in the repo. Should I create it as an empty placeholder, or is it managed externally?
5. **"All projects"** — how many projects are in scope for this handoff, and should I use the `cross_project` tools to read them?
6. **Go toolchain** — sandbox has no `go` binary. For any Go fix I make, are you OK with me invoking `nix run nixpkgs#go -- build ./...` to verify, or should I patch blind and you verify locally?
