// 問題: https://atcoder.jp/contests/abc140/tasks
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func abc140C() {
	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	n, _ := strconv.Atoi(strings.Fields(in)[0])
	in, _ = r.ReadString('\n')
	arr := strings.Fields(in)
	arr = append(arr, "100000")

	ans, _ := strconv.Atoi(arr[0])

	for i := 0; i < n - 1; i++ {
		b0, _ := strconv.Atoi(arr[i])
		b1, _ := strconv.Atoi(arr[i + 1])

		if b0 <= b1 {
			ans += b0
		} else {
			ans += b1
		}
	}

	fmt.Print(ans)
}
