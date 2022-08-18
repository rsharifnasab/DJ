#!/usr/bin/env bash

set -euo pipefail

readonly COMMAND="$1"

if [[ "$COMMAND" == "compile" ]]; then
    rm -rf out
    mkdir -p out
    python -m compileall -q .
elif [[ "$COMMAND" == "run" ]]; then
    inp_file="$2"
    actual_out="$3"
    python3 src/main.py <"$inp_file" >"$actual_out"
else
    echo "error in python lang"
fi
