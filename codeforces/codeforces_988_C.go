package codeforces

import (
	"fmt"
)

// https://codeforces.com/contest/2037/problem/C
func Codeforces_988_C() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)

		if n < 5 {
			fmt.Println("-1")
			continue
		}

		for i := n; i >= 6; i-- {
			if i%2 == 0 {
				fmt.Print(i, " ")
			}
		}
		fmt.Print("2 4 5 3 1 ")
		for i := 7; i <= n; i++ {
			if i%2 == 1 {
				fmt.Print(i, " ")
			}
		}
		fmt.Println("")
	}
}
