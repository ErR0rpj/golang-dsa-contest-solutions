package codeforces

import (
	"fmt"
)

// https://codeforces.com/contest/2031/problem/C
func Codeforces_987_C() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)

		if n%2 == 0 {
			for i := 1; i <= n; i++ {
				ans := i / 2
				if i%2 == 1 {
					ans++
				}
				fmt.Print(ans, " ")
			}
			fmt.Println("")
			continue
		} else {
			if n < 27 {
				fmt.Println("-1")
				continue
			} else {
				for i := 1; i <= n; i++ {
					if i == 1 || i == 10 || i == 26 {
						fmt.Print("1 ")
					} else if i == 11 || i == 27 {
						fmt.Print("6 ")
					} else {
						fmt.Print((i/2)+1, " ")
					}
				}
				fmt.Println("")
				continue
			}
		}
	}
}
