// 問題: https://atcoder.jp/contests/abc071/tasks
package main

import (
	"fmt"
	"strings"
)

func abc071B() {
	var alp = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var s string
	fmt.Scanf("%s", &s)
	l := strings.Split(s, "")

	m := map[string]int{}
	for i, ss := range l {
		m[ss] = i
	}

	for _, a := range alp {
		if _, ok := m[a]; !ok {
			fmt.Println(a)
			return
		}
	}

	fmt.Println("None")
}
