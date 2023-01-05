package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abc142C() {
	var n int
	fmt.Scan(&n)

	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	arr := strings.Fields(in)

	l := make([]int, n)
	for i := range l {
		a, _ := strconv.Atoi(arr[i])
		l[a - 1] = i + 1
	}

	for _, a := range l {
		fmt.Print(a, " ")
	}
}
