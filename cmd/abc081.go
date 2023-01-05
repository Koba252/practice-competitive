// 問題: https://atcoder.jp/contests/abc081/tasks
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func abc081B() {
	var n int
	fmt.Scanf("%d", &n)

	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	arr := strings.Fields(in)

	l := make([]int, n)
	for i, a := range arr {
		l[i], _ = strconv.Atoi(a)
	}

	var ans int
	for {
		for i, ll := range l {
			if ll % 2 != 0 {
				fmt.Println(ans)
				return
			}

			l[i] = ll / 2
		}
		ans++
	}
}
