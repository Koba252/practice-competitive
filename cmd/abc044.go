// 問題: https://atcoder.jp/contests/abc044/tasks
package main

import (
	"fmt"
)

func abc044B() {
	var w string
	fmt.Scanf("%s", &w)

	m := make(map[rune]int)
	for _, a := range w {
		if _, ok := m[a]; ok {
			m[a]++
		} else {
			m[a] = 1
		}
	}

	for _, v := range m {
		if v % 2 == 1 {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
