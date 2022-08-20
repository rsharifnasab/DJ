#!/usr/bin/env bash

set -euo pipefail

readonly COMMAND="$1"

if [[ "$COMMAND" == "compile" ]]; then
    rm -rf out
    mkdir -p out
    gcc -std=c11 -w -O2 -fomit-frame-pointer -lm -lpthread src/main.c -o out/a.out
elif [[ "$COMMAND" == "run" ]]; then
    inp_file="$2"
    actual_out="$3"
    ./out/a.out <"$inp_file" >"$actual_out"
else
    echo "error in java lang"
fi
