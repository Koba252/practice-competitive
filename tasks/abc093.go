// 問題: https://atcoder.jp/contests/abc093/tasks/abc093_b
package tasks

import (
	"fmt"
)

func ABC093B() {
	var a, b, k int
	fmt.Scanf("%d %d %d", &a, &b, &k)

	if b-a+1 <= k*2 {
		for i := a; i <= b; i++ {
			fmt.Println(i)
		}
		return
	}

	for i := 0; i < k; i++ {
		fmt.Println(a + i)
	}
	for i := 0; i < k; i++ {
		fmt.Println(b - k + 1 + i)
	}
}
