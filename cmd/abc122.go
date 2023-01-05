// 問題: https://atcoder.jp/contests/abc122/tasks
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func abc122B() {
	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	arr := strings.Split(in, "")

	var now, max int
	for _, a := range arr {
		if a == "A" || a == "C" || a == "G" || a == "T" {
			now++
		} else {
			if now > max {
				max = now
			}
			now = 0
		}
	}

	if now > max {
		max = now
	}

	fmt.Println(max)
}
