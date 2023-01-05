// 問題: https://atcoder.jp/contests/abc079/tasks
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func abc079C() {
	r := bufio.NewReader(os.Stdin)
	in, _ := r.ReadString('\n')
	in = strings.TrimSpace(in)
	arr := strings.Split(in, "")

	a, _ := strconv.Atoi(arr[0])
	b, _ := strconv.Atoi(arr[1])
	c, _ := strconv.Atoi(arr[2])
	d, _ := strconv.Atoi(arr[3])

	aarr := []int{a}

	var barr []int
	if b != 0 {
		barr = []int{b, -b}
	} else {
		barr = []int{b}
	}

	var carr []int
	if c != 0 {
		carr = []int{c, -c}
	} else {
		carr = []int{c}
	}

	var darr []int
	if d != 0 {
		darr = []int{d, -d}
	} else {
		darr = []int{d}
	}

	var total int
	for _, ar := range aarr {
		total += ar
		for _, br := range barr {
			total += br
			for _, cr := range carr {
				total += cr
				for _, dr := range darr {
					total += dr
					if total == 7 {
						fmt.Println(int2str([]int{ar, br, cr, dr}))
						return
					}

					total -= dr
				}
				total -= cr
			}
			total -= br
		}
		total -= ar
	}
}

func int2str(l []int) string {
	var s []string
	for i, n := range l {
		if i !=0 && n >= 0 {
			s = append(s, "+" + strconv.Itoa(n))
		} else {
			s = append(s, strconv.Itoa(n))
		}
	}

	return strings.Join(s, "") + "=7"
}
