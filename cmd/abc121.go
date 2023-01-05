// 問題: https://atcoder.jp/contests/abc121/tasks
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func abc121B() {
	var n, m , c int
	fmt.Scanf("%d %d %d", &n, &m, &c)

	r := bufio.NewReader(os.Stdin)
	b, _ := r.ReadString('\n')
	bArr := strings.Fields(b)
	var bl []int
	for _, bb := range bArr {
		bbb, _ := strconv.Atoi(bb)
		bl = append(bl, bbb)
	}

	var cnt int
	for i := 0; i < n; i++ {
		a, _ := r.ReadString('\n')
		aArr := strings.Fields(a)

		var sum int
		for j := 0; j < m; j++ {
			aa, _ := strconv.Atoi(aArr[j])
			sum += bl[j] * aa
		}

		if sum + c > 0 {
			cnt++
		}
	}

	fmt.Println(cnt)
}
