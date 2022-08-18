#!/bin/bash
set -euo pipefail

readonly COMMAND="$1" #"count" or "test"
readonly LANG="$2"    #
readonly TEST_NUMBER="${3:-0}"
readonly OUT_PRINT_LIMIT="1000000"

echerr() {
    echo "$@" 1>&2
}

test_count() {
    find "testgroup/in" -name "input*.txt" | wc -l
}

compare() {
    local ACTUAL_FILE="$1"
    local EXPECTED_FILE="$2"
    local actual
    local expected

    actual=$(cat "$ACTUAL_FILE" | sed 's/[[:space:]]/ /g')
    expected=$(cat "$EXPECTED_FILE" | sed 's/[[:space:]]/ /g')
    if [[ "$actual" == "$expected" ]]; then
        printf "pass"
    else
        printf "fail"
        echerr ""
        echerr "----------actual:---------"
        if (("${#actual}" < "$OUT_PRINT_LIMIT")); then
            echerr "$actual"
        else
            echerr "<<actual is too big to display>>"
        fi
        echerr "---------expected:--------"
        if (("${#expected}" < "$OUT_PRINT_LIMIT")); then
            echerr "$expected"
        else
            echerr "<<expected is too big to display>>"
        fi
        echerr "--------------------------"
    fi
}

find_file() {
    local NUMBER="$1"
    local POSTFIX="$2"
    local list
    list="$(find "testgroup/${POSTFIX}" -type f -name "*.txt" | sort)"
    lines="$(wc -l <<<"$list")"
    if [[ lines -lt NUMBER ]]; then
        echerr "invalid test number [$NUMBER]"
        exit 1
    fi
    echo "$list" | head -n "$NUMBER" | tail -1
}

run_test() {
    local expected_file
    local inp_file

    local actual_out="actual.txt"
    rm -f "$actual_out"

    expected_file=$(find_file "$TEST_NUMBER" "out")
    inp_file=$(find_file "$TEST_NUMBER" "in")
    test_name="$TEST_NUMBER"
    echerr "------- ${test_name} ------"

    run_log="$(run_code "$inp_file" "$actual_out")"
    echerr "----------<runlog>---------"
    echerr "$run_log"
    echerr "---------</runlog>---------"

    result="$(compare "$actual_out" "$expected_file")"
    printf "test[%s] - %s\n" "$test_name" "$result"
}

clean() {
    rm "actual.txt"
}

run_code() {
    inp_file="$1"
    actual_out="$2"
    "./lang/${LANG}.sh" run "$inp_file" "$actual_out"
}

compile_code() {
    "./lang/${LANG}.sh" compile
}

main() {
    cd -P -- "$(dirname -- "$0")"
    if [[ "compile" == "$COMMAND" ]]; then
        compile_code
    elif [[ "count" == "$COMMAND" ]]; then
        test_count
    elif [[ "test" == "$COMMAND" ]]; then
        run_test "$TEST_NUMBER"
    elif [[ "clean" == "$COMMAND" ]]; then
        clean
    fi
}

main
