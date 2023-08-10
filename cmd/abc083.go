// 問題: https://atcoder.jp/contests/abc083/tasks
package main

import (
	"fmt"
)

func abc083B() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	var ans int
	for i := 1; i <= n; i++ {
		a1, a2 := sumAll(i, 0)
		tmp := a1 + a2
		if tmp >= a && tmp <= b {
			ans += i
		}
	}

	fmt.Print(ans)
}

func sumAll(a, b int) (int, int) {
	a1 := a/10
	a2 := a%10
	if a1 < 10 {
		return 0, a1 + a2 + b
	}
	return sumAll(a1, a2 + b)
}
