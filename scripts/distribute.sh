#!/usr/bin/env bash
set -e
set -o pipefail
set -o nounset

./scripts/build.sh

cp -r ./examples/question-add ./bin
cp -r ./examples/sol-add-c/ ./bin
cp -r ./examples/judge-simple ./bin
(
    cd bin
    zip -q -r ./DJ-dist.zip -- ./*
)
