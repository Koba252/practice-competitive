// 問題: https://atcoder.jp/contests/abc063
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	abc063B()
}

func abc063B() {
	var s string
	fmt.Scanf("%s", &s)

	sl := strings.Split(s, "")

	sort.Strings(sl)

	for i := 1; i < len(sl); i++ {
		if sl[i] == sl[i - 1] {
			fmt.Println("no")
			return
		}
	}

	fmt.Println("yes")
}
