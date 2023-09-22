// 問題: https://atcoder.jp/contests/abc133/tasks
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"strconv"
)

func abc133B() {
	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	arr := strings.Fields(in)

	var n, d int
	n, _ = strconv.Atoi(arr[0])
	d, _ = strconv.Atoi(arr[1])

	l := make([][]int, n)
	for i := 0; i < n; i++ {
		in, _ = r.ReadString('\n')
		arr = strings.Fields(in)
		for j := 0; j < d; j++ {
			tmp, _ := strconv.Atoi(arr[j])
			l[i] = append(l[i], tmp)
		}
	}

	var ans int
	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			var sum float64
			for k := 0; k < d; k++ {
				sum += math.Pow(float64(l[i][k] - l[j][k]), 2)
			}

			sumSqrt := math.Sqrt(sum)
			if sumSqrt == float64(int(sumSqrt)) {
				ans++
			}
		}
	}

	fmt.Print(ans)
}
