#!/bin/bash

PORT=8080

echo "Echo server started on port $PORT"

while true; do
    mkfifo /tmp/pipe$$
    cat /tmp/pipe$$ | while IFS= read -r line; do
        echo "echo $line"
    done | nc -l -p $PORT -k > /tmp/pipe$$
    rm /tmp/pipe$$
done
