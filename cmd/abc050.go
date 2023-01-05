// 問題: https://atcoder.jp/contests/abc055/tasks
package main

import (
	"fmt"
)

func abc050B() {
	var n int
	fmt.Scanf("%d", &n)

	p := 1
	for i := 1; i <= n; i++ {
		p *= i
		p %= 1000000007
	}

	fmt.Println(p)
}
