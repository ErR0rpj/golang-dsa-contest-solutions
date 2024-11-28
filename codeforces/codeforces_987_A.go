package codeforces

import (
	"fmt"
	"main/templates"
)

// https://codeforces.com/contest/2031/problem/A
func Codeforces_987_A() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)

		height := make([]int, n)
		templates.InputArray(height, n)

		maxiRepCount, repCount, currentNo := 0, 0, 0
		for i := range height {
			if currentNo != height[i] {
				currentNo = height[i]
				repCount = 1
			} else {
				repCount++
			}

			if repCount > maxiRepCount {
				maxiRepCount = repCount
			}
		}

		fmt.Println(n - maxiRepCount)
	}
}
