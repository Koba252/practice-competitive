// 問題: https://atcoder.jp/contests/abc139/tasks
package main

import (
	"fmt"
)

// abc139B ABC139のB回答
func abc139B() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	var n int = (b - 1) / (a - 1)
	var m int = (b - 1) % (a - 1)
	if m == 0 {
		fmt.Println(n)
		return
	}
	fmt.Println(n + 1)
}

// abc139C ABC139のC回答
func abc139C() {
	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	n, _ := strconv.Atoi(strings.Fields(in)[0])

	in, _ = r.ReadString('\n')
	arrStr := strings.Fields(in)

	arr := make([]int, n)
	for i, a := range arrStr {
		arr[i], _ = strconv.Atoi(a)
	}

	var ans, tmp int
	for i := 0; i < n - 1; i++ {
		if arr[i] >= arr[i + 1] {
			tmp++
		} else {
			if tmp > ans {
				ans = tmp
			}
			tmp = 0
		}
	}

	if tmp > ans {
		ans = tmp
	}

	fmt.Println(ans)
}
