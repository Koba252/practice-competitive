package main

import (
	"fmt"
	"strings"
	"strconv"
)

func abc059B() {
	var a, b string
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)

	if a == b {
		fmt.Print("EQUAL")
		return
	}

	if len(a) > len(b) {
		fmt.Print("GREATER")
		return
	} else if len(a) < len(b) {
		fmt.Print("LESS")
		return
	}

	aArray := strings.Split(a, "")
	bArray := strings.Split(b, "")

	for i, aStr := range aArray {
		aInt, _ := strconv.Atoi(aStr)
		bInt, _ := strconv.Atoi(bArray[i])

		if aInt > bInt {
			fmt.Print("GREATER")
			return
		} else if aInt < bInt {
			fmt.Print("LESS")
			return
		}
	}

	fmt.Print("EQUAL")
	return
}
