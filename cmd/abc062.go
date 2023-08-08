// 問題: https://atcoder.jp/contests/abc062/tasks
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func abc062B() {
	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	arr := strings.Fields(in)

	h, _ := strconv.Atoi(arr[0])
	w, _ := strconv.Atoi(arr[1])

	ans := ""
	for j := 0; j < w + 2; j++ {
		ans += "#"
	}
	ans += "\n"

	for i := 0; i < h; i++ {
		in, _ = r.ReadString('\n')
		arr = strings.Fields(in)
		ans += "#"
		ans += strings.Join(arr, "")
		ans += "#\n"
	}

	for j := 0; j < w + 2; j++ {
		ans += "#"
	}
	ans += "\n"

	fmt.Println(ans)
}
