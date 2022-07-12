#!/bin/bash

set -e
set -o pipefail
set -u

run_test(){
    printf "test[%d]: pass\n" "$1"
    true
}


COMMAND="$1"  #"count" or "test"
if [[ "count" == "$COMMAND" ]]; then
    echo "5"
elif [[ "test" == "$COMMAND" ]]; then
    TEST_NUMBER="$2"
    case "$TEST_NUMBER" in
    ''|*[!0-9]*) 
        echo "number is not valid"
        exit 1
        ;;
    *) 
        run_test "$TEST_NUMBER"
        ;;
    esac
    
fi
