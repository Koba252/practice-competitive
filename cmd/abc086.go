// 問題: https://atcoder.jp/contests/abc086/tasks
package main

import (
	"fmt"
	"strconv"
)

func main() {
	abc086B()
}

func abc086B() {
	var a, b string
	fmt.Scanf("%s %s", &a, &b)

	s := a + b
	n, _ := strconv.Atoi(s)

	i := 1
	for {
		if i * i == n {
			fmt.Println("Yes")
			return
		}

		if i >= n/2 {
			fmt.Println("No")
			return
		}

		i++
	}
}
