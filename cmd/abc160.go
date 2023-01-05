// 問題: https://atcoder.jp/contests/abc160/tasks
package main

import (
	"fmt"
)

func abc160C() {
	var k, n int
	fmt.Scan(&k)
	fmt.Scan(&n)

	var a1 int
	fmt.Scan(&a1)
	b := a1

	var sum int
	var max int
	for i := 0; i < n - 1; i++ {
		var a int
		fmt.Scan(&a)

		diff := a - b
		sum += diff
		b = a

		if diff >= max {
			max = diff
		}
	}

	c := k + a1 - b
	if c >= max {
		fmt.Println(sum)
	} else {
		fmt.Println(sum + c - max)
	}
}
