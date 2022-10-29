// 問題: https://atcoder.jp/contests/abc132/tasks
package main

import (
	"fmt"
	"sort"
)

func main() {
	abc132C()
}

func abc132C() {
	var n int
	fmt.Scan(&n)

	l := make([]int, n)
	for i := range l {
		fmt.Scan(&l[i])
	}

	sort.Ints(l)

	a := l[n/2 - 1]
	b := l[n/2]

	fmt.Println(b - a)
}
