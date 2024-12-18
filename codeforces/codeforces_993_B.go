package codeforces

import "fmt"

func Codeforces_993_B() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var s string
		fmt.Scan(&s)

		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == 'q' {
				fmt.Print("p")
			} else if s[i] == 'p' {
				fmt.Print("q")
			} else {
				fmt.Print("w")
			}
		}
		fmt.Println("")
	}
}
