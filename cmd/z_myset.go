// 自分用
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mySet() {
	// 標準入力1
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	// 標準入力2
	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	arr := strings.Fields(in)
	// Fieldsを使わない場合は下記
	// in = strings.TrimSpace(in)
	// arr := strings.Split(in, " ")

	// 標準出力1
	fmt.Println(0)

	// 標準出力2
	fmt.Println(strings.Join(arr, ""))

	// 文字列→数値
	var str string = "10"
	strInt, _ := strconv.Atoi(str)
	fmt.Print(strInt)

	// 数値→文字列
	var num int = 20
	numStr := strconv.Itoa(num)
	fmt.Print(numStr)
}

// 絶対値
func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

// 最大公約数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

// nextPermutation 配列aを辞書順で次の配列に並び替える
func nextPermutation(a []int) bool {
	return permutationPandita(a, func(x, y int) bool { return x < y })
}

// prevPermutation 配列aを辞書順で前の配列に並び替える
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
	for p, q := i+1, len(a)-1; p < q; p, q = p+1, q-1 {
		a[p], a[q] = a[q], a[p]
	}
	return true
}
