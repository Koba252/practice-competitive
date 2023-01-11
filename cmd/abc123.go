// 問題: https://atcoder.jp/contests/abc123/tasks
package main

import (
	"fmt"
)

func abc123B() {
	l := make([]int, 5)
	var max int
	for i := 0; i < 5; i++ {
		var a int
		fmt.Scanf("%d", &a)
		l[i] = flat10ABC123(a)

		if flat10ABC123(a) - a > flat10ABC123(max) - max {
			max = a
		}
	}

	max -= flat10ABC123(max)

	for _, v := range l {
		max += v
	}

	fmt.Print(max)
}

func flat10ABC123(n int) int {
	if n % 10 == 0 {
		return n
	}
	return  10 * (n / 10 + 1)
}
