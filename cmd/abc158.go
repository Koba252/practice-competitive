// 問題: https://atcoder.jp/contests/abc158/tasks
package main

import (
	"fmt"
)

func main() {
	abc158B()
}

func abc158B() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	sho := n / (a + b)
	amari := n % (a + b)

	ans := sho * a
	if amari < a {
		ans += amari
	} else {
		ans += a
	}

	fmt.Println(ans)
}
