#!/usr/bin/env bash
set -e
set -o pipefail
set -o nounset

rm -rf bin
mkdir bin

#go clean -cache

go test "./..." # -count=1
go vet "./..."

#GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o bin/linux.out
# disable cgo_disable to use tree-sitter natives
GOOS=linux GOARCH=amd64                go build -ldflags="-w -s" -o bin/linux.out
du -h bin/linux.out

GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o bin/windows.exe
du -h bin/windows.exe

GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o bin/darwin.out
du -h bin/darwin.out
