// 問題: https://atcoder.jp/contests/abc060/tasks
package main

import (
	"fmt"
)

func abc060B() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	for i := 0; i <= b; i++ {
		if a*i % b == c {
			fmt.Println("YES")
			return
		}
	}

	fmt.Println("NO")
}
