// 自分用
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
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

	// 数値→文字列
	var num int = 20
	numStr := strconv.Itoa(num)
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
