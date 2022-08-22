#!/usr/bin/env bash

set -euo pipefail

readonly COMMAND="$1"

if [[ "$COMMAND" == "compile" ]]; then
    rm -rf out
    mkdir -p out
    python -m compileall -q .
elif [[ "$COMMAND" == "run" ]]; then
    src_file="$2"
    compiled_file="$3"
    (
        cd src
        python3 main.py -i "../${src_file}" -o "../${compiled_file}"
    )
else
    echo "error in python lang"
fi
