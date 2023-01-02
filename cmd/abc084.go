// 問題: https://atcoder.jp/contests/abc084/tasks
package main

import (
	"fmt"
	"strings"
)

func main() {
	abc084B()
}

func abc084B() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	var s string
	fmt.Scanf("%s", &s)

	sAry := strings.Split(s, "-")
	if len(sAry) != 2 || len(sAry[0]) != a || len(sAry[1]) != b {
		fmt.Println("No")
		return
	}

	fmt.Println("Yes")
}
