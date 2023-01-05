// 問題: https://atcoder.jp/contests/abc157/tasks
package main

import (
	"fmt"
)

func abc157B() {
	var a11, a12, a13, a21, a22, a23, a31, a32, a33 int
	fmt.Scanf("%d %d %d", &a11, &a12, &a13)
	fmt.Scanf("%d %d %d", &a21, &a22, &a23)
	fmt.Scanf("%d %d %d", &a31, &a32, &a33)

	l := [][]int{
		[]int{a11, a12, a13},
		[]int{a21, a22, a23},
		[]int{a31, a32, a33},
	}

	var n int
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		var b int
		fmt.Scanf("%d", &b)

		for j, ll := range l {
			for k, lll := range ll {
				if lll == b {
					l[j][k] = 0
					break
				}
			}
		}
	}

	if l[0][0] + l[0][1] + l[0][2] == 0 ||
		l[1][0] + l[1][1] + l[1][2] == 0 ||
		l[2][0] + l[2][1] + l[2][2] == 0 ||
		l[0][0] + l[1][0] + l[2][0] == 0 ||
		l[0][1] + l[1][1] + l[2][1] == 0 ||
		l[0][2] + l[1][2] + l[2][2] == 0 ||
		l[0][0] + l[1][1] + l[2][2] == 0 ||
		l[0][2] + l[1][1] + l[2][0] == 0 {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}
