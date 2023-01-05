package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func abc075B() {
	var h, w int
	fmt.Scanf("%d %d", &h, &w)

	r := bufio.NewReader(os.Stdin)

	var tmpAry []string
	for i := 0; i < w + 2; i++ {
		tmpAry = append(tmpAry, ".")
	}
	
	var sAry [][]string
	sAry = append(sAry, tmpAry)

	for i := 0; i < h; i++ {
		in, _ := r.ReadString('\n')
		in = strings.TrimSpace(in)
		inAry := []string{"."}
		inAry = append(inAry, strings.Split(in, "")...)
		inAry = append(inAry, ".")
		sAry = append(sAry, inAry)
	}

	sAry = append(sAry, tmpAry)

	for i := 1; i <= h; i++ {
		var tmpResult []string
		for j := 1; j <= w; j++ {
			var count int
			if sAry[i][j] == "#" {
				tmpResult = append(tmpResult, "#")
				continue
			}

			if sAry[i - 1][j - 1] == "#" {
				count++
			}
			if sAry[i - 1][j] == "#" {
				count++
			}
			if sAry[i - 1][j + 1] == "#" {
				count++
			}
			if sAry[i][j - 1] == "#" {
				count++
			}
			if sAry[i][j + 1] == "#" {
				count++
			}
			if sAry[i + 1][j - 1] == "#" {
				count++
			}
			if sAry[i + 1][j] == "#" {
				count++
			}
			if sAry[i + 1][j + 1] == "#" {
				count++
			}

			tmpResult = append(tmpResult, strconv.Itoa(count))
		}

		fmt.Println(strings.Join(tmpResult, ""))
	}
}
