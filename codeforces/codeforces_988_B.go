package codeforces

import (
	"fmt"
	"main/templates"
)

// https://codeforces.com/contest/2037/problem/B
func Codeforces_988_B() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var k int
		fmt.Scan(&k)
		a := make([]int, k)
		templates.InputArray(a, k)

		nm := k - 2
		var n int
		var m int
		count := make(map[int]int)

		for i := range a {
			count[a[i]]++
		}

		for i := range a {
			if nm%a[i] == 0 {
				count[a[i]]--
				if count[nm/a[i]] > 0 {
					n = a[i]
					m = nm / a[i]
					break
				} else {
					count[a[i]]++
				}
			}
		}

		fmt.Println(n, m)

	}
}
