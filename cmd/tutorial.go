package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	space2NewLine()
}

// 任意のn回分だけ入力データをそのまま出力する
func multiStdin() {
	r := bufio.NewReader(os.Stdin)

	in, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	in = strings.TrimSpace(in)

	n, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	
	for i := 0; i < n; i++ {
		s, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		s = strings.TrimSpace(s)
		fmt.Println(s)
	}
}

// スペース区切りの1行文字列を改行区切りで出力する
func space2NewLine() {
	r := bufio.NewReader(os.Stdin)

	in, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	in = strings.TrimSpace(in)

	// もしくは s := strings.Split(in, " ")
	s := strings.Fields(in)

	for _, out := range s {
		fmt.Println(out)
	}
}
