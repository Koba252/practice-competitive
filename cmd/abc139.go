// 問題: https://atcoder.jp/contests/abc139/tasks
package main

import (
	"fmt"
)

func main() {
	abc139B()
}

// abc139B ABC139のB回答
func abc139B() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	var n int = (b - 1) / (a - 1)
	var m int = (b - 1) % (a - 1)
	if m == 0 {
		fmt.Println(n)
		return
	}
	fmt.Println(n + 1)
}
