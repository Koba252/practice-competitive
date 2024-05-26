package main

import (
	"fmt"
	"os"

	"github.com/c6tower/practice-competitive/tasks"
)

var taskMap = map[string]func(){
	"ABC044B": tasks.ABC044B,
	"ABC046B": tasks.ABC046B,
	"ABC093B": tasks.ABC093B,
	"ABC158B": tasks.ABC158B,
	"ABC158C": tasks.ABC158C,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify the task name")
		os.Exit(1)
	}

	task := os.Args[1]
	if f, ok := taskMap[task]; ok {
		f()
	} else {
		fmt.Printf("Task not found: %s\n", task)
		os.Exit(1)
	}
}
