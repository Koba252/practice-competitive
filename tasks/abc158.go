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

	for ans := 1; ans <= 1000; ans++ {
		if 100*(ans+a) <= 108*ans && 100*(ans+a+1) > 108*ans && 100*(ans+b) <= 110*ans && 100*(ans+b+1) > 110*ans {
			fmt.Println(ans)
			return
		}
	}

	fmt.Println(-1)
}
