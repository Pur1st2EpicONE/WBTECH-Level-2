#!/bin/bash

CUT_OUTPUT="cut_output"
MY_CUT_OUTPUT="my_cut_output"

run_test() {

    local file="$1"

    cut -f 0 "$file" > "$CUT_OUTPUT" 2>&1
    ./cut -f 0 "$file" > "$MY_CUT_OUTPUT" 2>&1

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -f 0 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -f 0 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT" 

    cut -f 2 "$file" > "$CUT_OUTPUT"
    ./cut -f 2 "$file" > "$MY_CUT_OUTPUT"

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -f 2 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -f 2 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut -f 2- "$file" > "$CUT_OUTPUT" 2>&1
    ./cut -f 2- "$file" > "$MY_CUT_OUTPUT" 2>&1

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -f 2- test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -f 2- test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT" 

    cut -f -2 "$file" > "$CUT_OUTPUT" 2>&1
    ./cut -f -2 "$file" > "$MY_CUT_OUTPUT" 2>&1

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -f -2 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -f -2 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT" 

    cut -f 0-3 "$file" > "$CUT_OUTPUT" 2>&1
    ./cut -f 0-3 "$file" > "$MY_CUT_OUTPUT" 2>&1

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -f 0-3 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -f 0-3 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT" 

    cut -f 1-0 "$file" > "$CUT_OUTPUT" 2>&1
    ./cut -f 1-0 "$file" > "$MY_CUT_OUTPUT" 2>&1

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -f 1-0 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -f 1-0 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT" 

    cut -f -1-3 "$file" > "$CUT_OUTPUT" 2>&1
    ./cut -f -1-3 "$file" > "$MY_CUT_OUTPUT" 2>&1

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -f -1-3 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -f -1-3 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT" 

    cut -f 1,3-5 -s "$file" > "$CUT_OUTPUT"
    ./cut -f 1,3-5 -s "$file" > "$MY_CUT_OUTPUT"

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -f 1,3-5 -s test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -f 1,3-5 -s test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut -d ',' -f 2-3 "$file" > "$CUT_OUTPUT"
    ./cut -d ',' -f 2-3 "$file" > "$MY_CUT_OUTPUT"

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -d ',' -f 2-3 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -d ',' -f 2-3 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT" 

    cut -d ',' -f 2- "$file" > "$CUT_OUTPUT"
    ./cut -d ',' -f 2- "$file" > "$MY_CUT_OUTPUT"

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -d ',' -f 2- test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -d ',' -f 2- test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT" 

    cut -d ',' -f -3 "$file" > "$CUT_OUTPUT"
    ./cut -d ',' -f -3 "$file" > "$MY_CUT_OUTPUT"

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -d ',' -f -3 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -d ',' -f -3 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT" 

    cut -d ':' -f 2 "$file" > "$CUT_OUTPUT"
    ./cut -d ':' -f 2 "$file" > "$MY_CUT_OUTPUT"

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -d ':' -f 2 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -d ':' -f 2 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT" 

    cut -d ':' -f 1-100 "$file" > "$CUT_OUTPUT"
    ./cut -d ':' -f 1-100 "$file" > "$MY_CUT_OUTPUT"

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -d ':' -f 1-100 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -d ':' -f 1-100 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut -d '!' -f 1 "$file" > "$CUT_OUTPUT"
    ./cut -d '!' -f 1 "$file" > "$MY_CUT_OUTPUT"

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -d '!' -f 1 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -d '!' -f 1 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut -d ',' -f 10 "$file" > "$CUT_OUTPUT"
    ./cut -d ',' -f 10 "$file" > "$MY_CUT_OUTPUT"

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -d ',' -f 10 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -d ',' -f 10 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut -d '!' -f 1-2,4-5 "$file" > "$CUT_OUTPUT"
    ./cut -d '!' -f 1-2,4-5 "$file" > "$MY_CUT_OUTPUT"

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -d '!' -f 1-2,4-5 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -d '!' -f 1-2,4-5 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut -d '!' -f 0-2 "$file" > "$CUT_OUTPUT" 2>&1
    ./cut -d '!' -f 0-2 "$file" > "$MY_CUT_OUTPUT" 2>&1

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -d '!' -f 0-2 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -d '!' -f 0-2 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut -d ',' -f 3-3 "$file" > "$CUT_OUTPUT"
    ./cut -d ',' -f 3-3 "$file" > "$MY_CUT_OUTPUT"

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -d ',' -f 3-3 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -d ',' -f 3-3 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut -d ',' -f 1-3,3,3-5 -s "$file" > "$CUT_OUTPUT"
    ./cut -d ',' -f 1-3,3,3-5 -s "$file" > "$MY_CUT_OUTPUT"

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -d ',' -f 1-3,3,3-5 -s test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -d ',' -f 1-3,3,3-5 -s test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut -d '|' -f 2 "$file" > "$CUT_OUTPUT"
    ./cut -d '|' -f 2 "$file" > "$MY_CUT_OUTPUT"

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -d '|' -f 2 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: cut -d '|' -f 2 test_file_1.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut -d ',' -f 2 "$file" ./assets/test_file_2.txt > "$CUT_OUTPUT"
    ./cut -d ',' -f 2 "$file" ./assets/test_file_2.txt > "$MY_CUT_OUTPUT"

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -d ',' -f 2 test_file_1.txt test_file_2.txt"
    else
        echo "============================================"
        echo "Test failed: cut -d ',' -f 2 test_file_1.txt test_file_2.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut -d ':' -f 3 -s "$file" ./assets/test_file_2.txt > "$CUT_OUTPUT"
    ./cut -d ':' -f 3 -s "$file" ./assets/test_file_2.txt > "$MY_CUT_OUTPUT"

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: cut -d ':' -f 3 -s test_file_1.txt test_file_2.txt"
    else
        echo "============================================"
        echo "Test failed: cut -d ':' -f 3 -s test_file_1.txt test_file_2.txt"
        echo "Expected (cut):"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi

    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT" temp_test


    cut -d ',' -f --2 "$file" > "$CUT_OUTPUT" 2>&1
    ./cut -d ',' -f --2 "$file" > "$MY_CUT_OUTPUT" 2>&1

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: Invalid field range error"
    else
        echo "============================================"
        echo "Test failed: Invalid field range error"
        echo "Expected:"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut -d ',' -f a "$file" > "$CUT_OUTPUT" 2>&1
    ./cut -d ',' -f a "$file" > "$MY_CUT_OUTPUT" 2>&1

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: Invalid field value error"
    else
        echo "============================================"
        echo "Test failed: Invalid field value error"
        echo "Expected:"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut -d ',' -f 3-2 "$file" > "$CUT_OUTPUT" 2>&1
    ./cut -d ',' -f 3-2 "$file" > "$MY_CUT_OUTPUT" 2>&1

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: Invalid decreasing range error"
    else
        echo "============================================"
        echo "Test failed: Invalid decreasing range error"
        echo "Expected:"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut -d ',' -f 3-4 "NoSuchFile" > "$CUT_OUTPUT" 2>&1
    ./cut -d ',' -f 3-4 "NoSuchFile" > "$MY_CUT_OUTPUT" 2>&1

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: No such file or directory error"
    else
        echo "============================================"
        echo "Test failed: No such file or directory error"
        echo "Expected:"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut -d ',' -f 1, 3 "$file" > "$CUT_OUTPUT" 2>&1
    ./cut -d ',' -f 1, 3 "$file" > "$MY_CUT_OUTPUT" 2>&1

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: Fields are numbered from 1 error"
    else
        echo "============================================"
        echo "Test failed: Fields are numbered from 1 error"
        echo "Expected:"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

    cut > "$CUT_OUTPUT" 2>&1
    ./cut > "$MY_CUT_OUTPUT" 2>&1

    if diff -u "$CUT_OUTPUT" "$MY_CUT_OUTPUT"; then
        echo "Test passed: Must specify a list of fields error"
    else
        echo "============================================"
        echo "Test failed: Must specify a list of fields error"
        echo "Expected:"
        cat "$CUT_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_CUT_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$CUT_OUTPUT" "$MY_CUT_OUTPUT"

}

go build -o cut ./cmd/cut/main.go
run_test "./assets/test_file_1.txt"
