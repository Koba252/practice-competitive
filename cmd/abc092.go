// 問題: https://atcoder.jp/contests/abc092/tasks
package main

import (
	"fmt"
)

func abc092B() {
	var n, d, x int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d %d", &d, &x)

	ans := n + x
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)

		cnt := 1
		for {
			tmp := cnt * a + 1
			if tmp > d {
				break
			}

			ans++
			cnt++
		}
	}

	fmt.Println(ans)
}
