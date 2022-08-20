#!/usr/bin/env bash

set -euo pipefail

readonly COMMAND="$1"

if [[ "$COMMAND" == "compile" ]]; then
    rm -rf out
    mkdir -p out
    find . -name "*.java" >out/sources.txt
    javac -d out --release 11 -cp ".:lib/*" @out/sources.txt

elif [[ "$COMMAND" == "run" ]]; then
    inp_file="$2"
    actual_out="$3"
    java -cp "out:lib/*" Main <"$inp_file" >"$actual_out"
else
    echo "error in java lang"
fi
