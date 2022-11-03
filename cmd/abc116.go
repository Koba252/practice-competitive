// 問題: https://atcoder.jp/contests/abc116/tasks
package main

import (
	"fmt"
)

func main() {
	abc116B()
}

func abc116B() {
	var s int
	fmt.Scanf("%d", &s)

	if s == 1 || s == 2 || s == 4 {
		fmt.Println(4)
		return
	}

	var cnt int
	for s != 4 {
		if s % 2 == 0 {
			s = s / 2
			cnt++
		} else {
			s = 3 * s + 1
			cnt++
		}
	}

	fmt.Println(cnt + 4)
}
