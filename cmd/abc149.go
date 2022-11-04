// 問題: https://atcoder.jp/contests/abc149/tasks
package main

import (
	"fmt"
)

func main() {
	abc149C()
}

func abc149C() {
	var x int
	fmt.Scanf("%d", &x)

	for {
		if isPrime(x) {
			fmt.Println(x)
			return
		}

		x++
	}
}

func isPrime(n int) bool {
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
