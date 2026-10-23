# scripts/ — PowerShell Module Architecture

## Overview

The `run.ps1` dispatcher (≤130 lines) imports all modules from this directory and routes commands to the appropriate function. Each `.psm1` module is self-contained with PowerShell help documentation.

## Module Reference

| Module | Lines | Purpose |
|---|---|---|
| `DashboardTheme.psm1` | 158 | Theme detection, ANSI color initialization |
| `DashboardBoxes.psm1` | 229 | Box primitives, progress bars, section renderers |
| `DashboardPhases.psm1` | 145 | Phase registry & summary box |
| `DashboardCoverage.psm1` | 218 | Coverage table & diff rendering |
| `DashboardUI.psm1` | 55 | Thin facade: composite `Write-Dashboard` |
| `Utilities.psm1` | 105 | Console output helpers, line filtering/merging |
| `ErrorParser.psm1` | 195 | Build/runtime error extraction & classification |
| `TestLogWriter.psm1` | 214 | Parse Go test output → structured log files |
| `TestRunner.psm1` | 276 | Go test execution, build checks, git ops |
| `CoveragePreChecks.psm1` | 145 | safetest, autofix, bracecheck, in-pkg import validation |
| `check-inpkg-imports.ps1` | — | Lint: forbid heavy framework imports in in-package tests |
| `check-safetest-boundaries.ps1` | — | Lint: safeTest closure boundaries |
| `check-integrated-regressions.ps1` | — | API-drift regression scanner |
| `CoverageCompileCheck.psm1` | 273 | Compile checks & per-file split recovery |
| `CoverageProfileMerger.psm1` | 120 | Profile merging & missing profile detection |
| `CoverageReport.psm1` | 554 | Report generation (TXT, JSON, HTML, AI) |
| `CoverageRunner.psm1` | 298 | TC orchestrator |
| `PackageCoverage.psm1` | 171 | TCP single-package coverage |
| `BuildTools.psm1` | 137 | Build, run, format, vet, tidy, clean |
| `GoConvey.psm1` | 57 | Browser-based GoConvey test runner |
| `PreCommitCheck.psm1` | 181 | Pre-commit API mismatch checker |
| `Help.psm1` | 140 | Help display, fail log viewer |

## Dependency Graph

```
DashboardTheme           (standalone)
DashboardBoxes           (→ DashboardTheme)
DashboardPhases          (→ DashboardTheme, DashboardBoxes)
DashboardCoverage        (→ DashboardTheme, DashboardBoxes)
DashboardUI              (→ all Dashboard*)
Utilities                (standalone, optional DashboardUI)
ErrorParser              (standalone)
TestLogWriter            (→ Utilities)
TestRunner               (→ Utilities, TestLogWriter)
CoveragePreChecks        (→ Utilities, DashboardPhases)
CoverageCompileCheck     (→ Utilities, ErrorParser)
CoverageProfileMerger    (standalone)
CoverageReport           (→ Utilities, ErrorParser, DashboardCoverage)
CoverageRunner           (→ all Coverage*, Utilities, TestLogWriter, TestRunner)
PackageCoverage          (→ CoveragePreChecks, CoverageReport, TestRunner)
BuildTools               (→ Utilities)
GoConvey                 (standalone)
PreCommitCheck           (→ Utilities, CoveragePreChecks, DashboardPhases)
Help                     (→ Utilities, TestLogWriter, TestRunner)
```

## How the Dispatch Works

1. `run.ps1` defines `param($Command, [string[]]$ExtraArgs)`
2. Imports all `.psm1` modules in dependency order via a loop
3. `switch ($Command.ToLower())` routes to the matching function

## Adding a New Command

1. Create or extend a module with the new function
2. Export via `Export-ModuleMember -Function @('Your-Function')`
3. Add to `$moduleOrder` in `run.ps1` if it's a new module
4. Add a switch case in `run.ps1`
5. Update `Show-Help` in `Help.psm1`

## Module Loading

All modules load via `Import-Module -Force -DisableNameChecking` in dependency order. Cross-module calls work because all modules share the same session scope. DashboardUI functions are guarded with `Get-Command ... -ErrorAction SilentlyContinue`.
