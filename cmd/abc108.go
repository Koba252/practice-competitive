// 問題: https://atcoder.jp/contests/abc108/tasks
package main

import (
	"fmt"
)

func abc108B() {
	var x1, y1, x2, y2, x3, y3, x4, y4 int
	fmt.Scanf("%d %d %d %d", &x1, &y1, &x2, &y2)

	xDiff := x1 - x2
	yDiff := y1 - y2

	x3 = x2 + yDiff
	y3 = y2 - xDiff
	x4 = x1 + yDiff
	y4 = y1 - xDiff

	fmt.Println(x3, y3, x4, y4)
}
