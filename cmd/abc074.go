// 問題: https://atcoder.jp/contests/abc074
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	abc074B()
}

func abc074B() {
	var n, k int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &k)

	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	arr := strings.Fields(in)
	
	var sum int
	for _, a := range arr {
		x, _ := strconv.Atoi(a)

		if k - x > x {
			sum += x * 2
		} else {
			sum += (k - x) * 2
		}
	}

	fmt.Println(sum)
}
