// 問題: https://atcoder.jp/contests/abc109/tasks
package main

import (
	"fmt"
)

func abc109B() {
	var n int
	fmt.Scanf("%d", &n)

	used := map[string]int{}
	var a, b string
	fmt.Scanf("%s", &a)
	used[a] = 999
	for i := 1; i < n; i++ {
		fmt.Scanf("%s", &b)
		if used[b] != 0 {
			fmt.Print("No")
			return
		}

		if fmt.Sprintf("%c", b[0]) != fmt.Sprintf("%c", a[len(a)-1]) {
			fmt.Print("No")
			return
		}

		used[b] = i
		a = b
	}

	fmt.Print("Yes")
}
