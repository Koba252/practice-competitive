// 問題: https://atcoder.jp/contests/abc090/tasks
package main

import (
	"fmt"
	"strings"
	"strconv"
)

func abc090B() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	var ans int
	for i := a; i <= b; i++ {
		s := strconv.Itoa(i)
		array := strings.Split(s, "")
		len := len(array)
		half := len / 2
		hasDif := false
		for j := 0; j < half; j++ {
			if array[j] != array[len-1-j] {
				hasDif = true
				break
			}
		}

		if !hasDif {
			ans++
		}
	}

	fmt.Println(ans)
}
