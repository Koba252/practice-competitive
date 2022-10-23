// 問題: https://atcoder.jp/contests/abc088
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"strconv"
)

func main() {
	abc088B()
}

func abc088B() {
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

	sort.Sort(sort.Reverse(sort.IntSlice(l)))

	var a, b int
	for i := 0; i < n; i++ {
		if i % 2 == 0 {
			a += l[i]
		} else {
			b += l[i]
		}
	}

	fmt.Println(a - b)
}
