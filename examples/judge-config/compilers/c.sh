#!/usr/bin/env bash

set -o errexit
set -o nounset

command=@1

filename=@2

if [ $command = "compile" ]; then
    echo "compile"
    gcc -std=c11    -w -O2 -fomit-frame-pointer -lm -lpthread  $filename -o compiled/a.exe
elif [ "$command" = "run" ]; then
    echo "run"
    ./compild/a.exe
else 
    echo "error"
fi


