#!/bin/bash

GREP_OUTPUT="grep_output"
MY_GREP_OUTPUT="my_grep_output"
MY_GREP_OUTPUT_COLOURLESS="my_grep_output_colourless"

run_test() {

    local file="$1"
    
    grep door "$file" > "$GREP_OUTPUT"
    ./grep door "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep d..r "$file" > "$GREP_OUTPUT"
    ./grep d..r "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep d..r test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep d..r test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep door -A 2 "$file" > "$GREP_OUTPUT"
    ./grep door -A 2 "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -A 2 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -A 2 test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep door -A 3 "$file" > "$GREP_OUTPUT"
    ./grep door -A 3 "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -A 3 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -A 3 test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep door -B 2 "$file" > "$GREP_OUTPUT"
    ./grep door -B 2 "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -B 2 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -B 2 test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep door -B 3 "$file" > "$GREP_OUTPUT"
    ./grep door -B 3 "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -B 3 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -B 3 test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep door -C 2 "$file" > "$GREP_OUTPUT"
    ./grep door -C 2 "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -C 2 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -C 2 test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep door -C 3 "$file" > "$GREP_OUTPUT"
    ./grep door -C 3 "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -C 3 test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -C 3 test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep door -c "$file" > "$GREP_OUTPUT"
    ./grep door -c "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -c test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -c test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep Door -i "$file" > "$GREP_OUTPUT"
    ./grep Door -i "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep Door -i test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep Door -i test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep door -v "$file" > "$GREP_OUTPUT"
    ./grep door -v "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -v test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -v test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep d..r -F "$file" > "$GREP_OUTPUT"
    ./grep d..r -F "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep d..r -F test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep d..r -F test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep ...r -n "$file" > "$GREP_OUTPUT"
    ./grep ...r -n "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep ...r -n test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep ...r -n test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep door -nc "$file" > "$GREP_OUTPUT"
    ./grep door -nc "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -nc test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -nc test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep door -ni "$file" > "$GREP_OUTPUT"
    ./grep door -ni "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -ni test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -ni test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep [dr] -nv "$file" > "$GREP_OUTPUT"
    ./grep [dr] -nv "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep [dr] -nv test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep [dr] -nv test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep door -nF "$file" > "$GREP_OUTPUT"
    ./grep door -nF "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -nF test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -nF test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"    

    grep DOOR -ci "$file" > "$GREP_OUTPUT"
    ./grep DOOR -ci "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep DOOR -ci test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep DOOR -ci test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"   

    grep ...$ -cv "$file" > "$GREP_OUTPUT"
    ./grep ...$ -cv "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep ...$ -cv test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep ...$ -cv test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS" 

    grep door -cF "$file" > "$GREP_OUTPUT"
    ./grep door -cF "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -cF test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -cF test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS" 

    grep door -iv "$file" > "$GREP_OUTPUT"
    ./grep door -iv "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -iv test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -iv test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS" 

    grep ^..r -iF "$file" > "$GREP_OUTPUT"
    ./grep ^..r -iF "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep ^..r -iF test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep ^..r -iF test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS" 

    grep door -vF "$file" > "$GREP_OUTPUT"
    ./grep door -vF "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -vF test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -vF test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS" 



    grep door -A 2 -i -n "$file" > "$GREP_OUTPUT"
    ./grep door -A 2 -i -n "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -A 2 -v -n test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -A 2 -v -n test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"  

    grep door -B 4 -F -c "$file" > "$GREP_OUTPUT"
    ./grep door -B 4 -F -c "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -B 4 -F -c test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -B 4 -F -c test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS" 

    grep door -C 6 -n -i "$file" > "$GREP_OUTPUT"
    ./grep door -C 6 -n -i "$file" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door -C 6 -n -i test_file_1.txt"
    else
        echo "============================================"
        echo "Test failed: grep door -C 6 -n -i test_file_1.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep door "$file" ""./assets/test_file_2.txt"" > "$GREP_OUTPUT"
    ./grep door "$file" ""./assets/test_file_2.txt"" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep door test_file_1.txt test_file_2.txt"
    else
        echo "============================================"
        echo "Test failed: grep door test_file_1.txt test_file_2.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep d..r "$file" "./assets/test_file_2.txt" > "$GREP_OUTPUT"
    ./grep d..r "$file" "./assets/test_file_2.txt" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep d..r test_file_1.txt test_file_2.txt"
    else
        echo "============================================"
        echo "Test failed: grep d..r test_file_1.txt test_file_2.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep ^on -ni "$file" "./assets/test_file_2.txt" > "$GREP_OUTPUT"
    ./grep ^on -ni "$file" "./assets/test_file_2.txt" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep ^on -ni test_file_1.txt test_file_2.txt"
    else
        echo "============================================"
        echo "Test failed: grep ^on -ni test_file_1.txt test_file_2.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep D..r -cv "$file" "./assets/test_file_2.txt" > "$GREP_OUTPUT"
    ./grep D..r -cv "$file" "./assets/test_file_2.txt" > "$MY_GREP_OUTPUT"
    sed -r 's/\x1B\[[0-9;]*[mK]//g' "$MY_GREP_OUTPUT" > "$MY_GREP_OUTPUT_COLOURLESS"

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"; then
        echo "Test passed: grep D..r -cv test_file_1.txt test_file_2.txt"
    else
        echo "============================================"
        echo "Test failed: grep D..r -cv test_file_1.txt test_file_2.txt"
        echo "Expected (grep):"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT_COLOURLESS"
        echo "============================================"
    fi

    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT" "$MY_GREP_OUTPUT_COLOURLESS"

    grep > "$GREP_OUTPUT" 2>&1
    ./grep > "$MY_GREP_OUTPUT" 2>&1  

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT"; then
        echo "Test passed: No arguments error"
    else
        echo "============================================"
        echo "Test failed: No arguments error"
        echo "Expected:"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT"

    grep door "NoSuchFile" > "$GREP_OUTPUT" 2>&1
    ./grep door "NoSuchFile" > "$MY_GREP_OUTPUT" 2>&1

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT"; then
        echo "Test passed: No such file or directory error"
    else
        echo "============================================"
        echo "Test failed: No such file or directory error"
        echo "Expected:"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT"

    grep door ./assets/test_file_1.txt -A -1 > "$GREP_OUTPUT" 2>&1
    ./grep door ./assets/test_file_1.txt -A -1 > "$MY_GREP_OUTPUT" 2>&1  

    if diff -u "$GREP_OUTPUT" "$MY_GREP_OUTPUT"; then
        echo "Test passed: Invalid context length argument error"
    else
        echo "============================================"
        echo "Test failed: Invalid context length argument error"
        echo "Expected:"
        cat "$GREP_OUTPUT"
        echo "--------------------------------------------"
        echo "Got:"
        cat "$MY_GREP_OUTPUT"
        echo "============================================"
    fi
    
    rm -f "$GREP_OUTPUT" "$MY_GREP_OUTPUT"

}

go build -o grep ./cmd/grep/main.go
run_test "./assets/test_file_1.txt"
