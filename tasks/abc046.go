// 問題: https://atcoder.jp/contests/abc046/tasks
package tasks

import (
	"fmt"
)

func ABC046B() {
	// 標準入力1
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	r := k
	for i := 1; i < n; i++ {
		r *= (k - 1)
	}

	fmt.Println(r)
}
