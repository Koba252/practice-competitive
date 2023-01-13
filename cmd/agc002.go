// 問題: https://atcoder.jp/contests/agc002/tasks
package main

import (
	"fmt"
)

func agc002A() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	if a <= 0 && b >= 0 {
		fmt.Println("Zero")
		return
	}

	if a > 0 {
		fmt.Println("Positive")
		return
	}

	if (b - a) % 2 == 0 {
		fmt.Println("Negative")
		return
	}

	fmt.Println("Positive")
}
