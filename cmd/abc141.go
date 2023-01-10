// 問題: https://atcoder.jp/contests/abc141/tasks
package main

import (
	"fmt"
)

func abc141C() {
	var n, k, q int
	fmt.Scanf("%d %d %d", &n, &k, &q)

	aMap := make(map[int]int, n)
	for i := 0; i < q; i++ {
		var a int
		fmt.Scanf("%d", &a)

		aMap[a]++
	}

	for i := 1; i <= n; i++ {
		if k - q + aMap[i] > 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
