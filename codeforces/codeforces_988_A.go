package codeforces

import (
	"fmt"
	"main/templates"
)

func Codeforces_988_A() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)
		a := make([]int, n)
		templates.InputArray(a, n)

		countOfA := make([]int, n+1)
		for _, i := range a {
			countOfA[i]++
		}

		ans := 0
		for _, i := range countOfA {
			if i > 0 {
				if i%2 == 1 {
					i = i - 1
				}
				ans += (i / 2)
			}
		}

		fmt.Println(ans)

	}
}
