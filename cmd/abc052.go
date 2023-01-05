// 問題: https://atcoder.jp/contests/abc052/tasks
package main

import (
	"fmt"
	"strings"
)

func abc052B() {
	var n int
	var s string
	fmt.Scanf("%d", &n)
	fmt.Scanf("%s", &s)

	l := strings.Split(s, "")

	var max, cnt int
	for i := 0; i < n; i++ {
		if l[i] == "I" {
			cnt++
		} else {
			cnt--
		}

		if cnt > max {
			max = cnt
		}
	}

	fmt.Println(max)
}
