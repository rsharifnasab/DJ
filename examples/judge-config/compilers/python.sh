#!/usr/bin/env bash

set -o errexit
set -o nounset

command=$1

filename=$2

if [ "$command" = "compile" ]; then
    echo "compile"

    #TODO: compile python!

elif [ "$command" = "run" ]; then
    echo "run"

    python3 "$filename"

else
    echo "error"
fi


