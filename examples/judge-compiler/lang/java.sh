#!/usr/bin/env bash

set -euo pipefail

readonly COMMAND="$1"

if [[ "$COMMAND" == "compile" ]]; then
    rm -rf out
    mkdir -p out
    find . -name "*.java" >out/sources.txt
    javac -d out --release 11 -cp ".:lib/*" @out/sources.txt

elif [[ "$COMMAND" == "run" ]]; then
    src_file="$2"
    compiled_file="$3"
    java -cp "out:lib/*" Main -i "$src_file" -o "$compiled_file"
else
    echo "error in java lang"
fi
