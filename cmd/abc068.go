// 問題: https://atcoder.jp/contests/abc068/tasks
package main

import (
	"fmt"
)

func abc068B() {
	var n int
	fmt.Scan(&n)

	var result int
	var tmp int
	for i := 1; i <= n; i++ {
		var cnt int
		a := i
		for j := 1; j < 7; j++ {
			if a % 2 == 0 {
				a /= 2
				cnt++
			} else {
				break
			}
		}

		if cnt >= tmp {
			tmp = cnt
			result = i
		}
	}

	fmt.Println(result)
}
