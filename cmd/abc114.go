// 問題: https://atcoder.jp/contests/abc114/tasks
package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	abc114B()
}

func abc114B() {
	var s string
	fmt.Scanf("%s", &s)

	sl := strings.Split(s, "")

	ans := 753
	for i := 0; i < len(sl) - 2; i++ {
		as := sl[i] + sl[i + 1] + sl[i + 2]
		a, _ := strconv.Atoi(as)

		diff := abs(a - 753)
		if diff < ans {
			ans = diff
		}
	}

	fmt.Println(ans)
}

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}
