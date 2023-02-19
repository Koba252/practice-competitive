// 問題: https://atcoder.jp/contests/abc152/tasks
package main

import (
	"fmt"
)

func abc152C() {
	var n, max int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &max)

	ans := 1
	min := max
	for i := 1; i < n; i++ {
		var p int
		fmt.Scanf("%d", &p)

		if p > max {
			max = p
		} else if p < min {
			min = p
			ans++
		}
	}

	fmt.Println(ans)
}
