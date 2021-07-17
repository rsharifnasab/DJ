#!/usr/bin/env bash

set -o errexit
set -o nounset

command=@1

filename=@2

if [ $command = "compile" ]; then
    echo "compile"
    g++ -std=c++17  -w -O2 -fomit-frame-pointer -lm  -pthread  $filename -o compiled/a.exe
elif [ "$command" = "run" ]; then
    echo "run"
    ./compild/a.exe
else 
    echo "error"
fi


