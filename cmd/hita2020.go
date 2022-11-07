// 問題: https://atcoder.jp/contests/hitachi2020/tasks
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
	hita2020B()
}

func hita2020B() {
	var a, b, m int
	fmt.Scanf("%d %d %d", &a, &b, &m)

	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	arr := strings.Fields(in)
	al := make([]int, a)
	for i, s := range arr {
		ss, _ := strconv.Atoi(s)
		al[i] = ss
	}

	in, _ = r.ReadString('\n')
	arr = strings.Fields(in)
	bl := make([]int, b)
	for i, s := range arr {
		ss, _ := strconv.Atoi(s)
		bl[i] = ss
	}

	ans := al[0] + bl[0]
	for i := 0; i < m; i++ {
		in, _ = r.ReadString('\n')
		arr = strings.Fields(in)
		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])
		c, _ := strconv.Atoi(arr[2])

		tmp := al[x - 1] + bl[y - 1] - c
		if tmp < ans {
			ans = tmp
		}
	}

	sort.Ints(al)
	sort.Ints(bl)

	if al[0] + bl[0] < ans {
		ans = al[0] + bl[0]
	}

	fmt.Println(ans)
}
