#!/bin/bash

WorkDir="$PWD/.."
BinDir=$WorkDir/bin

echo ""
echo " ---- [Start] Running all in docker [Start]-----"
echo ""
echo "Work dir     : $WorkDir"
echo "Binaries dir : $BinDir"

# docker run --rm -it -v "PWD":/usr/src/myapp -v "$GOPATH":/go -w /usr/src/myapp golang:1.17.8

rm -rf "$BinDir"/results &&
  mkdir -p "$BinDir"/results &&
  chown -R root:root "$BinDir" &&
  chmod -R 777 "$BinDir" &&
  ls -la "$BinDir" &&
  echo "" &&
  docker run --rm -it -v "$WorkDir":/usr/src/myapp -v "$GOPATH":/go -w /usr/src/myapp golang:1.17.8 bash -c '
./bin/cli-linux-amd64 2>&1 | tee bin/results/linux-amd64.out && cat bin/results/linux-amd64.out
' &&
  echo "Running complete" &&
  echo "" &&
  echo "Output" &&
  echo "" &&
  echo "" &&
  echo "ls -la $BinDir/results" &&
  cat "$BinDir/results/linux-amd64.out" &&
  ls -lah "$BinDir/results" &&
  echo "$\"Path\":" &&
  echo "export PATH=\$PATH:\"$BinDir\"" &&
  echo "running : ${BinDir}/cli-linux-amd64" &&
  echo "" &&
  echo " ---- [End] Running in docker [end]-----" &&
  echo ""
