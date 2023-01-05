package main

import (
	// "bufio"
	"fmt"
	"math"
	// "os"
	// "sort"
	// "strconv"
	// "strings"
)

// abc086C ABC086Cの回答
func abc086C() {
	var n int
	fmt.Scanf("%d", &n)

	var t, x, y, pt, px, py int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d", &t, &x, &y)

		sum := int(math.Abs(float64(x - px)) + math.Abs(float64(y - py)))
		dt := t - pt
		if sum > dt || (dt - sum) % 2 == 1 {
			fmt.Println("No")
			return
		}

		pt = t
		px = x
		py = y
	}

	fmt.Println("Yes")
}

// // abc049C ABC049Cの回答
// func abc049C() {
// 	var s string
// 	fmt.Scanf("%s", &s)

// 	for {
// 		tmpS := s
// 		s = strings.TrimSuffix(s, "dreamer")
// 		s = strings.TrimSuffix(s, "dream")
// 		s = strings.TrimSuffix(s, "eraser")
// 		s = strings.TrimSuffix(s, "erase")

// 		if len(s) == 0 {
// 			fmt.Println("YES")
// 			return
// 		}

// 		if len(s) == len(tmpS) {
// 			fmt.Println("NO")
// 			return
// 		}
// 	}
// }

// // abc085C ABC085Cの回答
// func abc085C() {
// 	var n, y int
// 	fmt.Scanf("%d %d", &n, &y)

// 	a := -1
// 	b := -1
// 	c := -1
// 	isBrreak := false
// 	for i := 0; i <= n; i++ {
// 		if isBrreak {
// 			break
// 		}
// 		for j := 0; j <= n - i; j++ {
// 			k := n - i - j
// 			if (i * 10000 + j * 5000 + k * 1000) == y {
// 				a = i
// 				b = j
// 				c = k
// 				break
// 			}
// 		}
// 	}

// 	fmt.Printf("%d %d %d\n", a, b, c)
// }

// // abc085B() ABC085Bの回答
// func abc085B() {
// 	var n int
// 	fmt.Scanf("%d", &n)

// 	var l []int
// 	for i := 0; i < n; i++ {
// 		var d int
// 		fmt.Scanf("%d", &d)

// 		var hasValue bool
// 		for _, a := range l {
// 			if d == a {
// 				hasValue = true
// 				break
// 			}
// 		}

// 		if !hasValue {
// 			l = append(l, d)
// 		}
// 	}

// 	fmt.Println(len(l))
// }

// // abc088B ABC088Bの回答
// func abc088B() {
// 	var n int
// 	fmt.Scanf("%d", &n)

// 	r := bufio.NewReader(os.Stdin)
// 	s, _ := r.ReadString('\n')
// 	sAry := strings.Fields(s)
// 	var aAry []int
// 	for _, sa := range sAry {
// 		a, _ := strconv.Atoi(sa)
// 		aAry = append(aAry, a)
// 	}

// 	sort.Sort(sort.Reverse(sort.IntSlice(aAry)))
// 	var sumA int
// 	var sumB int
// 	for i, a := range aAry {
// 		if (i % 2) == 0 {
// 			sumA += a
// 		} else {
// 			sumB += a
// 		}
// 	}

// 	fmt.Println(sumA - sumB)
// }

// // abc083B ABC083Bの回答
// func abc083B() {
// 	var n, a, b int
// 	fmt.Scanf("%d %d %d", &n, &a, &b)

// 	var sum int
// 	for i := 1; i <= n; i++ {
// 		iStr := strconv.Itoa(i)
// 		iStrAry := strings.Split(iStr, "")

// 		var tmpSum int
// 		for _, s := range iStrAry {
// 			sInt, _ := strconv.Atoi(s)
// 			tmpSum += sInt
// 		}

// 		if tmpSum >= a && tmpSum <= b {
// 			sum += i
// 		}
// 	}

// 	fmt.Println(sum)
// }

// // abc087B ABC087Bの回答
// func abc087B() {
// 	var a, b, c, x int
// 	fmt.Scanf("%d", &a)
// 	fmt.Scanf("%d", &b)
// 	fmt.Scanf("%d", &c)
// 	fmt.Scanf("%d", &x)


// 	var count int
// 	for i := 0; i <= a; i++ {
// 		for j := 0; j <= b; j++ {
// 			for k := 0; k <= c; k++ {
// 				sum := (500 * i) + (100 * j) + (50 * k)
// 				if sum == x {
// 					count++
// 				}
// 			}
// 		}
// 	}

// 	fmt.Println(count)
// }

// // abc081B ABC081Bの回答
// func abc081B() {
// 	r := bufio.NewReader(os.Stdin)

// 	nStr, err := r.ReadString('\n')
// 	if err != nil {
// 		panic(err)
// 	}
// 	nStr = strings.TrimSpace(nStr)

// 	n, err := strconv.Atoi(nStr)
// 	if err != nil {
// 		panic(err)
// 	}

// 	in, err := r.ReadString('\n')
// 	if err != nil {
// 		panic(err)
// 	}

// 	inAry := strings.Fields(in)

// 	var count int

// 	for i := 0; i < n; i++ {
// 		a, err := strconv.Atoi(inAry[i])
// 		if err != nil {
// 			panic(err)
// 		}

// 		tmpCount := 0
// 		for (a % 2) == 0 {
// 			a = a / 2
// 			tmpCount++
// 		}

// 		if i == 0 || tmpCount < count {
// 			count = tmpCount
// 		}
// 	}

// 	fmt.Println(count)
// }

// // abc081A ABC081Aの回答
// func abc081A() {
// 	var s string
// 	fmt.Scanf("%s", &s)

// 	var i int
// 	sAry := strings.Split(s, "")
// 	for _, sa := range sAry {
// 		if sa == "1" {
// 			i++
// 		}
// 	}
// 	fmt.Println(i)
// }

// // abc086A ABC086Aの回答
// func abc086A() {
// 	var a, b int
// 	if _, err := fmt.Scanf("%d", &a); err != nil {
// 		panic(err)
// 	}

// 	if _, err := fmt.Scanf("%d", &b); err != nil {
// 		panic(err)
// 	}

// 	if a * b % 2 == 1 {
// 		fmt.Println("Odd")
// 	} else {
// 		fmt.Println("Even")
// 	}
// }

// // beg01 PracticeA回答
// func beg01() {
// 	var a, b, c int
// 	var s string

// 	fmt.Scanf("%d", &a)
// 	fmt.Scanf("%d %d", &b, &c)
// 	fmt.Scanf("%s", &s)

// 	fmt.Printf("%d %s\n", a + b + c, s)
// }

// func smiple() {
// 	r := bufio.NewReader(os.Stdin)
// 	w := bufio.NewWriter(os.Stdout)
// 	defer w.Flush()

// 	var s string
// 	fmt.Fscan(r, &s)

// 	fmt.Fprintln(w, s)
// }
