#!/bin/bash
# Usage: ./run.sh <filename> <task>
# Example: ./run.sh abc044 B
# This will run the golang function ABC044B() with the input from `in.txt` and output to `out.txt`

task=$(echo $1 | tr '[:lower:]' '[:upper:]')$2

# Build the executable
go build -o tmp_executable cmd/main.go

# Run the executable with input and output redirection
cat in.txt | ./tmp_executable $task > out.txt

# Remove the temporary executable
rm tmp_executable
