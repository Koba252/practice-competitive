package main

import (
	"fmt"
	"sort"
)

func abc134C() {
	var n int
	fmt.Scanf("%d", &n)

	l1 := make([]int, n)
	l2 := make([]int, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d", &a)
		l1[i] = a
		l2[i] = a
	}

	sort.Ints(l2)

	for i, _ := range l1 {
		if l1[i] != l2[n-1] {
			fmt.Println(l2[n-1])
		} else {
			fmt.Println(l2[n-2])
		}
	}
}
