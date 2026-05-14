$WorkDir = (Get-Location).Path
$BinDir = Join-Path $WorkDir "bin"

Write-Host ""
Write-Host " ---- [Start] Running all in docker [Start]-----"
Write-Host ""
Write-Host "Work dir     : $WorkDir"
Write-Host "Binaries dir : $BinDir"

# docker run --rm -it -v "$PWD":/usr/src/myapp -v "$GOPATH":/go -w /usr/src/myapp golang:1.17.8

Remove-Item -Recurse -Force "$BinDir/results"
New-Item -Path "$BinDir/results" -ItemType Directory | Out-Null
Get-ChildItem -Path "$BinDir" | ForEach-Object { $_ | Set-ItemProperty -Name IsReadOnly -Value $false }
Get-ChildItem -Path "$BinDir" | ForEach-Object { $_ | Set-ItemProperty -Name Attributes -Value 'Directory' }
Get-ChildItem -Path "$BinDir" | Format-Table -Property Name, Attributes
Write-Host ""
docker run --rm -it -v "$WorkDir":/usr/src/myapp -v "$Env:GOPATH":/go -w /usr/src/myapp golang:1.17.8 bash -c '
./bin/cli-linux-amd64 2>&1 | tee bin/results/linux-amd64.out; cat bin/results/linux-amd64.out
'
Write-Host "Running complete"
Write-Host ""
Write-Host "Output"
Write-Host ""
Write-Host ""
Write-Host "ls -la $BinDir/results"
Get-Content "$BinDir/results/linux-amd64.out"
Get-ChildItem -Path "$BinDir/results" | Format-Table -Property Name, Length
Write-Host "$(\"Path\"):"
Write-Host "export PATH=\$PATH:\"$BinDir\""
Write-Host "running : ${BinDir}/cli-linux-amd64"
Write-Host ""
Write-Host " ---- [End] Running in docker [end]-----"
Write-Host ""
