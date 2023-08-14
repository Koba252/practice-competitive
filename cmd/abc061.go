// 問題: https://atcoder.jp/contests/abc061/tasks
package main

import (
	"fmt"
)

func abc061B() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	ans := make([]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		ans[a-1] += 1
		ans[b-1] += 1
	}

	for _, a := range ans {
		fmt.Println(a)
	}
}
