// 問題: https://atcoder.jp/contests/agc024/tasks
package main

import (
	"fmt"
	"math"
)

func agc024A() {
	var a, b, c, k int
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &k)
	
	var ans int
	if k%2 == 0 {
		ans = a - b
	} else {
		ans = b - a
	}

	max := int(math.Pow10(18))
	if ans > max  || ans < -max {
		fmt.Println("Unfair")
		return
	}
	fmt.Println(ans)
}
