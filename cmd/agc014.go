// 問題: https://atcoder.jp/contests/agc014/tasks
package main

import (
	"fmt"
)

func agc014A() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	if a % 2 == 1 || b % 2 == 1 || c % 2 == 1 {
		fmt.Println(0)
		return
	}

	if a == b && b == c {
		fmt.Println(-1)
		return
	}

	var cnt int
	for a % 2 == 0 && b % 2 == 0 && c % 2 == 0 {
		aa := a
		bb := b
		cc := c
		a = bb / 2 + cc / 2
		b = cc / 2 + aa / 2
		c = aa / 2 + bb / 2

		cnt++
	}

	fmt.Println(cnt)
}
