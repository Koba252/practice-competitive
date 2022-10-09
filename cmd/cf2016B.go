// 問題: https://atcoder.jp/contests/code-festival-2016-qualb/tasks
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	var s string
	fmt.Scanf("%s", &s)
	s = strings.TrimSpace(s)
	arr := strings.Split(s, "")

	limit := a + b
	sum := 0
	sumB := 0
	for i := 0; i < n; i++ {
		if sum >= limit {
			for j := i; j < n; j++ {
				fmt.Println("No")
			}
			return
		}

		if arr[i] == "c" {
			fmt.Println("No")
			continue
		}

		if arr[i] == "b" && sumB < b {
			fmt.Println("Yes")
			sum++
			sumB++
			continue
		}

		if arr[i] == "a" {
			fmt.Println("Yes")
			sum++
			continue
		}

		fmt.Println("No")
	}
}
