#!/bin/bash

SORT_OUTPUT="sort_output"
MY_SORT_OUTPUT="my_sort_output"

run_test() {
    local file="$1"
    
    echo "Testing file: $file"
    
    sort -n "$file" > "$SORT_OUTPUT"
    go run *.go -n "$file" > "$MY_SORT_OUTPUT"

    if diff -u "$SORT_OUTPUT" "$MY_SORT_OUTPUT"; then
        echo "Test passed: -n"
    else
        echo "============================================"
        echo "Test failed: -n"
        echo "Expected (sort -n):"
        cat "$SORT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_SORT_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$SORT_OUTPUT" "$MY_SORT_OUTPUT"


    sort "$file" > "$SORT_OUTPUT"
    go run *.go "$file" > "$MY_SORT_OUTPUT"

    if diff -u "$SORT_OUTPUT" "$MY_SORT_OUTPUT"; then
        echo "Test passed: no flags"
    else
        echo "============================================"
        echo "Test failed: no flags"
        echo "Expected (sort):"
        cat "$SORT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_SORT_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$SORT_OUTPUT" "$MY_SORT_OUTPUT"

    sort -r "$file" > "$SORT_OUTPUT"
    go run *.go -r "$file" > "$MY_SORT_OUTPUT"

    if diff -u "$SORT_OUTPUT" "$MY_SORT_OUTPUT"; then
        echo "Test passed: -r"
    else
        echo "============================================"
        echo "Test failed: -r"
        echo "Expected (sort):"
        cat "$SORT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_SORT_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$SORT_OUTPUT" "$MY_SORT_OUTPUT"

    
    sort -rn "$file" > "$SORT_OUTPUT"
    go run *.go -rn "$file" > "$MY_SORT_OUTPUT"

    if diff -u "$SORT_OUTPUT" "$MY_SORT_OUTPUT"; then
        echo "Test passed: -rn"
    else
        echo "============================================"
        echo "Test failed: -rn"
        echo "Expected (sort):"
        cat "$SORT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_SORT_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$SORT_OUTPUT" "$MY_SORT_OUTPUT"

    sort -nu "$file" > "$SORT_OUTPUT"
    go run *.go -nu "$file" > "$MY_SORT_OUTPUT"

    if diff -u "$SORT_OUTPUT" "$MY_SORT_OUTPUT"; then
        echo "Test passed: -nu"
    else
        echo "============================================"
        echo "Test failed: -nu"
        echo "Expected:"
        cat "$SORT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_SORT_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$SORT_OUTPUT" "$MY_SORT_OUTPUT"

    sort -rnu "$file" > "$SORT_OUTPUT"
    go run *.go -rnu "$file" > "$MY_SORT_OUTPUT"

    if diff -u "$SORT_OUTPUT" "$MY_SORT_OUTPUT"; then
        echo "Test passed: -rnu"
    else
        echo "============================================"
        echo "Test failed: -rnu"
        echo "Expected:"
        cat "$SORT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_SORT_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$SORT_OUTPUT" "$MY_SORT_OUTPUT"

    sort -k 2 "test3" > "$SORT_OUTPUT"
    go run *.go "test3" -k 2 > "$MY_SORT_OUTPUT"

    if diff -u "$SORT_OUTPUT" "$MY_SORT_OUTPUT"; then
        echo "Test passed: -k 2"
    else
        echo "============================================"
        echo "Test failed: -k 2"
        echo "Expected:"
        cat "$SORT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_SORT_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$SORT_OUTPUT" "$MY_SORT_OUTPUT"

}

run_test "test"