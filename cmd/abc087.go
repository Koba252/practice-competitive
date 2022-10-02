// 問題: https://atcoder.jp/contests/abc087/tasks
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	abc087C()
}

func abc087C() {
	var n int
	fmt.Scanf("%d", &n)

	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	arr1 := strings.Fields(in)

	in2, _ := r.ReadString('\n')
	arr2 := strings.Fields(in2)

	var rslt int
	var sum1 int
	for i := 0; i < n; i++ {
		a1, _ := strconv.Atoi(arr1[i])
		sum1 += a1

		sum := sum1
		for j := i; j < n; j++ {
			a2, _ := strconv.Atoi(arr2[j])
			sum += a2
		}

		if sum > rslt {
			rslt = sum
		}
	}

	fmt.Println(rslt)
}
