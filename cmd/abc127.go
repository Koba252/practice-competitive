// 問題: https://atcoder.jp/contests/abc127/tasks
package main

import (
	"fmt"
)

func abc127C() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	min, max := 0, 100000
	for i := 0; i < m; i++ {
		var l, r int
		fmt.Scanf("%d %d", &l, &r)
		if l > min {
			min = l
		}
		if r < max {
			max = r
		}
	}

	ans := max - min + 1
	if ans < 0 {
		ans = 0
	}
	fmt.Println(ans)
}
