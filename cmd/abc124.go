// 問題: https://atcoder.jp/contests/abc124/tasks
package main

import (
	"fmt"
)

func abc124C() {
	var s string
	fmt.Scanf("%s", &s)

	var p, q int
	for i := 0; i < len(s); i++ {
		if (i % 2 == 0 && s[i] == '0') || (i % 2 == 1 && s[i] == '1') {
			p++
		} else if (i % 2 == 0 && s[i] == '1') || (i % 2 == 1 && s[i] == '0') {
			q++
		}
	}

	if p < q {
		fmt.Println(p)
	} else {
		fmt.Println(q)
	}
}
