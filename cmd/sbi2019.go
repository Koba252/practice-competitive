// 問題: https://atcoder.jp/contests/sumitrust2019/tasks
package main

import (
	"fmt"
)

func sbi2019B() {
	var n int
	fmt.Scanf("%d", &n)

	x := float64(n) / 1.08
	if x > float64(int(x)) {
		x += 1
	}

	out := int(x)

	if out * 108 >= n * 100 && out * 108 < 100 * (n + 1) {
		fmt.Println(out)
		return
	}

	fmt.Println(":(")
}
