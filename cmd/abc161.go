// 問題: https://atcoder.jp/contests/abc161/tasks
package main

import (
	"fmt"
)

func main() {
	abc161C()
}

func abc161C() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	n = n % k
	if n <= 1 {
		fmt.Println(n)
		return
	}

	if n < k - n {
		fmt.Println(n)
		return
	}

	fmt.Println(k - n)
}
