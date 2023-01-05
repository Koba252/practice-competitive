// 問題: https://atcoder.jp/contests/abc153/tasks
package main

import (
	"fmt"
)

func abc153D() {
	var h int
	fmt.Scanf("%d", &h)

	var ans int = 1
	var monsterNum int = 1
	for {
		h /= 2
		if h == 0 {
			break
		}

		monsterNum *= 2
		ans += monsterNum
	}

	fmt.Println(ans)
}
