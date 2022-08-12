#!/usr/bin/env bash
set -e
set -o pipefail
set -o nounset

rm -rf bin
mkdir bin

go test "./..." -count=1
go vet "./..."


GOOS=linux   GOARCH=amd64 go build -ldflags="-w -s"  -o bin/linux.out
du -h bin/linux.out

GOOS=windows GOARCH=amd64 go build -ldflags="-w -s"  -o bin/windows.exe
du -h bin/windows.exe

GOOS=darwin  GOARCH=amd64 go build -ldflags="-w -s"  -o bin/darwin.out
du -h bin/darwin.out

cp -r ./examples/judge-compiler-* ./bin
cp -r ./examples/question-compiler ./bin
cp -r ./examples/COMPILER_README.* ./bin
(
cd bin
zip -q -r ./compiler-phase3.zip -- ./*

)
