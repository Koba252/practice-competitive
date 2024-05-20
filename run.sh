#!/bin/bash
# Usage: ./run.sh <filename> <task>
# Example: ./run.sh abc044 B
# This will run the golang function abc044B() with the input from `in.txt` and output to `out.txt`

file="cmd/${1}.go"
task=$1$2
cat in.txt | go run cmd/main.go $file $task > out.txt
