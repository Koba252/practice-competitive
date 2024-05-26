// 問題: https://atcoder.jp/contests/abc158/tasks
package tasks

import (
	"fmt"
)

func ABC158B() {
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

func ABC158C() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	for i := 1; i <= 10000; i++ {
		var aa int = i * 8 / 100
		var bb int = i * 10 / 100
		if aa == a && bb == b {
			fmt.Println(i)
			return
		}
	}

	fmt.Println(-1)
}
