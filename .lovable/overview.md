# .lovable/ Overview

This folder is the project's institutional memory. It is co-maintained by the
user and the AI agent. The agent reads it on every session to recover context.

## Layout

```
.lovable/
├── overview.md           # This file
├── plan.md               # Active roadmap (single file, see ## Completed at bottom)
├── prompt.md             # Index of reusable prompts
├── prompts/              # Reusable prompt definitions
├── strictly-avoid.md     # Hard prohibitions
├── suggestions.md        # All suggestions (active + implemented)
├── cicd-index.md         # CI/CD issues index
├── cicd-issues/          # One file per CI/CD issue
├── pending-issues/       # Unresolved bugs
├── solved-issues/        # Resolved bugs (with Solution / Learning sections)
└── memory/
    ├── index.md          # Index of all memory files (MUST list every file)
    ├── 01-project-overview.md
    ├── workflow/         # Phase / state tracking
    ├── decisions/        # Architectural decisions
    ├── conventions/      # Project conventions
    ├── blockers/         # Active blockers needing user input
    └── avoid/            # Detail behind strictly-avoid.md entries
```

## Conventions

- All filenames lowercase + hyphen-separated, `XX-` numeric prefix where ordering matters.
- Single-file aggregates: `plan.md`, `suggestions.md`, `strictly-avoid.md`.
- Never delete entries; mark complete and keep them.
- Every file under `memory/` MUST appear in `memory/index.md`.

To checkpoint a session: run the prompt at `prompts/01-write-memory.md`.
