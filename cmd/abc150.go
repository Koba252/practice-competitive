// 問題: https://atcoder.jp/contests/abc150/tasks
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	abc150C()
}

func abc150C() {
	var n int
	fmt.Scanf("%d", &n)

	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	arr := strings.Fields(in)
	ps := make([]int, n)
	for i, _ := range ps {
		ps[i], _ = strconv.Atoi(arr[i])
	}

	in, _ = r.ReadString('\n')
	arr = strings.Fields(in)
	qs := make([]int, n)
	for i, _ := range qs {
		qs[i], _ = strconv.Atoi(arr[i])
	}

	base := make([]int, n)
	for i := range base {
		base[i] = i + 1
	}

	var ansP, ansQ int
	i := 1
	for {
		if isEqualSlice(base, ps) {
			ansP = i
		}

		if isEqualSlice(base, qs) {
			ansQ = i
		}

		if ansP != 0 && ansQ != 0 {
			break
		}

		nextPermutation(base)
		i++
	}

	ans := ansP - ansQ

	if ans < 0 {
		ans = -ans
	}

	fmt.Println(ans)
}

func nextPermutation(a []int) bool {
	return permutationPandita(a, func(x, y int) bool { return x < y })
}

func prevPermutation(a []int) bool {
	return permutationPandita(a, func(x, y int) bool { return x > y })
}

func permutationPandita(a []int, less func(x, y int) bool) bool {
	i := len(a) - 2
	for 0 <= i && !less(a[i], a[i+1]) {
		i--
	}

	if i == -1 {
		return false
	}

	j := i + 1
	for k := j; k < len(a); k++ {
		if less(a[i], a[k]) && !less(a[j], a[k]) {
			j = k
		}
	}

	a[i], a[j] = a[j], a[i]
	for p, q := i + 1, len(a) - 1; p < q; p, q = p + 1, q - 1 {
		a[p], a[q] = a[q], a[p]
	}
	return true
}

func isEqualSlice(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
