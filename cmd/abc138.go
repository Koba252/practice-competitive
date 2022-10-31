// 問題: https://atcoder.jp/contests/abc138/tasks
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
	abc138C()
}

func abc138C() {
	var n int
	fmt.Scanf("%d", &n)

	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	arr := strings.Fields(in)

	l := make([]float64, n)
	for i := range arr {
		a, _ := strconv.Atoi(arr[i])
		l[i] = float64(a)
	}

	sort.Float64s(l)

	for i := 1; i < n; i++ {
		l[i] = (l[i - 1] + l[i]) / 2
	}

	fmt.Println(l[n - 1])
}
