// 問題: https://atcoder.jp/contests/abc113/tasks
package tasks

import "fmt"

func ABC113B() {
	var N, T, A, H1 int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d %d", &T, &A)
	fmt.Scanf("%d", &H1)
	ans := 1
	diff0 := abs(1000*A - (1000*T - 6*H1))
	for i := 2; i < N+1; i++ {
		var Hi int
		fmt.Scanf("%d", &Hi)
		diff := abs(1000*A - (1000*T - 6*Hi))
		if diff < diff0 {
			ans = i
			diff0 = diff
		}
	}

	fmt.Println(ans)
}

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}
