#!/usr/bin/env pwsh
<#
.SYNOPSIS
    Lint check: in-package _test.go files must not import heavy test frameworks.
.DESCRIPTION
    Scans all _test.go files that live inside source packages (i.e., NOT under
    tests/ or scripts/) for imports of heavy test framework packages that cause
    [setup failed] errors during instrumented coverage runs (-coverpkg=./...).

    Forbidden import prefixes:
      - coretests/       (args, coretestcases, results, etc.)
      - github.com/smartystreets/goconvey
      - github.com/stretchr/testify

    Exit 0 = clean, Exit 1 = violations found.
#>

$ErrorActionPreference = "Stop"

# ── Configuration ─────────────────────────────────────────────────────────────

$forbiddenPatterns = @(
    '"[^"]*coretests/'
    '"github\.com/smartystreets/goconvey'
    '"github\.com/stretchr/testify'
)

# Directories that are allowed to use heavy imports (test-only directories)
$allowedDirs = @(
    'tests/'
    'scripts/'
    'coretests/'
)

# ── Scan ──────────────────────────────────────────────────────────────────────

$repoRoot = (Get-Location).Path
$issues = 0

# Find all _test.go files
$allTestFiles = Get-ChildItem -Path $repoRoot -Filter "*_test.go" -Recurse -File |
    Sort-Object FullName

foreach ($file in $allTestFiles) {
    $rel = $file.FullName.Replace($repoRoot, "").TrimStart([char]'\', [char]'/') -replace '\\', '/'

    # Skip files in allowed directories
    $skip = $false
    foreach ($allowed in $allowedDirs) {
        if ($rel.StartsWith($allowed)) {
            $skip = $true
            break
        }
    }
    if ($skip) { continue }

    # Read file and check imports
    $lines = Get-Content -LiteralPath $file.FullName
    $inImportBlock = $false

    for ($i = 0; $i -lt $lines.Count; $i++) {
        $line = $lines[$i].Trim()

        # Detect import block start
        if ($line -match '^import\s*\(') {
            $inImportBlock = $true
            continue
        }

        # Detect import block end
        if ($inImportBlock -and $line -eq ')') {
            $inImportBlock = $false
            continue
        }

        # Single-line import
        if ($line -match '^import\s+"') {
            foreach ($pattern in $forbiddenPatterns) {
                if ($line -match $pattern) {
                    $importPath = if ($line -match '"([^"]+)"') { $Matches[1] } else { $line }
                    Write-Host "  ${rel}:$($i + 1): forbidden import '$importPath' in in-package test"
                    $issues++
                }
            }
            continue
        }

        # Lines inside import block
        if ($inImportBlock) {
            foreach ($pattern in $forbiddenPatterns) {
                if ($line -match $pattern) {
                    $importPath = if ($line -match '"([^"]+)"') { $Matches[1] } else { $line }
                    Write-Host "  ${rel}:$($i + 1): forbidden import '$importPath' in in-package test"
                    $issues++
                }
            }
        }
    }
}

# ── Result ────────────────────────────────────────────────────────────────────

if ($issues -ne 0) {
    Write-Host ""
    Write-Host "✗ Found $issues forbidden import(s) in in-package test files."
    Write-Host "  Fix: use only 'testing' + stdlib, or move tests to tests/integratedtests/."
    exit 1
}

Write-Host "✓ All in-package test files use lightweight imports only."
exit 0
