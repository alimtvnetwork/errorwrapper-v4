#!/usr/bin/env pwsh
# Splits each *_test.go in corestrtests into its own subfolder (with shared support files),
# compiles each subfolder independently in parallel, and produces a pass/fail report.
# Usage: ./scripts/split-and-check-corestrtests.ps1

$ErrorActionPreference = "Stop"
$src = "tests/integratedtests/corestrtests"

if (-not (Test-Path $src)) {
    Write-Host "Directory $src not found." -ForegroundColor Red
    exit 1
}

# Collect test files, helper test files, and non-test support files
$testFiles = Get-ChildItem -LiteralPath $src -Filter "*_test.go" -File |
    Where-Object { $_.Name -notlike "*helper*" } | Sort-Object Name
$helperTestFiles = Get-ChildItem -LiteralPath $src -Filter "*helper*_test.go" -File | Sort-Object Name
$supportFiles = Get-ChildItem -LiteralPath $src -Filter "*.go" -File |
    Where-Object { $_.Name -notlike "*_test.go" } | Sort-Object Name

Write-Host "Found $($testFiles.Count) test files, $($helperTestFiles.Count) helper files, $($supportFiles.Count) support files."
Write-Host "Creating subfolders..."

# Create a subfolder per test file, copy it + all support files
foreach ($tf in $testFiles) {
    $folderName = $tf.BaseName -replace '_test$', ''
    $dest = Join-Path $src $folderName
    if (-not (Test-Path $dest)) { New-Item -ItemType Directory -Path $dest -Force | Out-Null }
    Copy-Item -LiteralPath $tf.FullName -Destination (Join-Path $dest $tf.Name) -Force
    foreach ($sf in $supportFiles) {
        Copy-Item -LiteralPath $sf.FullName -Destination (Join-Path $dest $sf.Name) -Force
    }
    foreach ($hf in $helperTestFiles) {
        Copy-Item -LiteralPath $hf.FullName -Destination (Join-Path $dest $hf.Name) -Force
    }
}

Write-Host "Created $($testFiles.Count) subfolders. Running parallel build checks..."

# Gather subfolders
$subfolders = Get-ChildItem -LiteralPath $src -Directory | Sort-Object Name

# Run go test -count=1 -run ^$ for each subfolder in parallel
$jobs = @()
foreach ($dir in $subfolders) {
    $pkg = "./tests/integratedtests/corestrtests/$($dir.Name)/"
    $jobs += Start-Job -ScriptBlock {
        param($pkg, $name)
        $result = & go test -count=1 -run '^$' $pkg 2>&1
        $rc = $LASTEXITCODE
        [PSCustomObject]@{
            Name     = $name
            ExitCode = $rc
            Output   = ($result | Out-String).Trim()
        }
    } -ArgumentList $pkg, $dir.Name
}

Write-Host "Waiting for $($jobs.Count) jobs..."
$results = $jobs | Wait-Job | Receive-Job
$jobs | Remove-Job -Force

$pass = @()
$fail = @()

foreach ($r in $results) {
    if ($r.ExitCode -eq 0) {
        $pass += $r.Name
    } else {
        $fail += $r
    }
}

$pass = $pass | Sort-Object
$fail = $fail | Sort-Object { $_.Name }

Write-Host ""
Write-Host "✓ PASS: $($pass.Count) subfolders" -ForegroundColor Green
Write-Host "✗ FAIL: $($fail.Count) subfolders" -ForegroundColor Red
Write-Host ""

foreach ($f in $fail) {
    Write-Host "  ✗ $($f.Name):" -ForegroundColor Red
    $lines = $f.Output -split "`n"
    foreach ($line in ($lines | Select-Object -First 8)) {
        Write-Host "    $line"
    }
    Write-Host ""
}

# Write machine-readable report
$report = @{
    pass = $pass
    fail = $fail | ForEach-Object { @{ name = $_.Name; error = $_.Output } }
}
$reportPath = "data/split-check-report.json"
if (-not (Test-Path "data")) { New-Item -ItemType Directory -Path "data" -Force | Out-Null }
$report | ConvertTo-Json -Depth 3 | Set-Content -Path $reportPath -Encoding UTF8
Write-Host "Report written to $reportPath"

# Cleanup: remove subfolders
Write-Host ""
Write-Host "Cleaning up subfolders..."
foreach ($dir in $subfolders) {
    Remove-Item -LiteralPath $dir.FullName -Recurse -Force
}
Write-Host "Done. Subfolders removed."
