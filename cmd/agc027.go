// 問題: https://atcoder.jp/contests/agc027/tasks
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
	agc027A()
}

func agc027A() {
	var n, x int
	fmt.Scanf("%d %d", &n, &x)

	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	arr := strings.Fields(in)

	l := make([]int, n)
	for i := range l {
		l[i], _ = strconv.Atoi(arr[i])
	}

	sort.Ints(l)

	var ans int
	for i := 0; i < n - 1; i++ {
		if l[i] > x {
			fmt.Println(ans)
			return
		}

		x -= l[i]
		ans++
	}

	if l[n - 1] == x {
		ans++
	}

	fmt.Println(ans)
}
