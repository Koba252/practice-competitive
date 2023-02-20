// 問題: https://atcoder.jp/contests/abc118/tasks
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func abc118B() {
	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	arr := strings.Fields(in)
	n, _ := strconv.Atoi(arr[0])

	in, _ = r.ReadString('\n')
	arr = strings.Fields(in)
	dict := make(map[string]int)
	for i := 1; i < len(arr); i++ {
		dict[arr[i]] = 1
	}


	for i := 1; i < n; i++ {
		in, _ = r.ReadString('\n')
		arr = strings.Fields(in)

		for j := 1; j < len(arr); j++ {
			if _, ok:= dict[arr[j]]; ok {
				dict[arr[j]]++
			}
		}
	}

	var ans int
	for _, v := range dict {
		if v == n {
			ans++
		}
	}

	fmt.Println(ans)
}
