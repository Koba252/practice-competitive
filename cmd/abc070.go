// 問題: https://atcoder.jp/contests/abc070/tasks
package main

import (
	"fmt"
)

func main() {
	abc070B()
}

func abc070B() {
	var a, b, c, d int
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	if a >= d || b < c {
		fmt.Println(0)
		return
	}

	if a <= c {
		if b <= d {
			fmt.Println(b - c)
			return
		} else {
			fmt.Println(d - c)
			return
		}
	} else {
		if b <= d {
			fmt.Println(b - a)
			return
		} else {
			fmt.Println(d - a)
			return
		}
	}
}
