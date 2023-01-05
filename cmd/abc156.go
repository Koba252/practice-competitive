// 問題: https://atcoder.jp/contests/abc156/tasks
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"strconv"
)

// ABC156のC回答
func abc156C() {
	var n int
	fmt.Scanf("%d", &n)

	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	arr := strings.Fields(in)
	var l []int
	for _, a := range arr {
		aa, _ := strconv.Atoi(a)
		l = append(l, aa)
	}

	sort.Ints(l)

	var newL []int
	for p := l[0]; p <= l[n - 1]; p++ {
		var tmp int
		for _, ll := range l {
			tmp += (p - ll) * (p - ll)
		}

		newL = append(newL, tmp)
	}

	sort.Ints(newL)

	fmt.Println(newL[0])
}
