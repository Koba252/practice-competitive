// 問題: https://atcoder.jp/contests/abc094/tasks
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func abc094B() {
	var n, m, x int
	fmt.Scanf("%d %d %d", &n, &m, &x)

	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	arr := strings.Fields(in)

	var s, l int
	for _, as := range arr {
		a, _ := strconv.Atoi(as)

		if a < x {
			s++
		} else {
			l++
		}
	}

	if s < l {
		fmt.Println(s)
	} else {
		fmt.Println(l)
	}
}
