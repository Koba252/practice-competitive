// 問題: https://atcoder.jp/contests/panasonic2020
package main

import (
	"fmt"
)

func main() {
	pana2020B()
}

func pana2020B() {
	var h, w int
	fmt.Scanf("%d %d", &h, &w)

	if h == 1 || w == 1 {
		fmt.Println(1)
		return
	}

	n := h * w
	if n % 2 == 0 {
		fmt.Println(w * h / 2)
	} else {
		fmt.Println(w * h / 2 + 1)
	}
}
