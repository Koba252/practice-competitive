package main

import (
	"fmt"
	"os"
)

var tasks = map[string]func(){
	"abc044B": abc044B,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify the task name")
		os.Exit(1)
	}

	task := os.Args[1]
	if f, ok := tasks[task]; ok {
		f()
	} else {
		fmt.Printf("Task not found: %s\n", task)
		os.Exit(1)
	}
}
