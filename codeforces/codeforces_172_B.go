package codeforces

import (
	"fmt"
)

func Codeforces_172_B() {

	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)

		arr := make([]int, n+1)
		singulars := 0
		score := 0
		totalColors := 0
		for i := 0; i < n; i++ {
			var now int
			fmt.Scan(&now)
			arr[now]++
			if arr[now] == 1 {
				totalColors++
				singulars++
			} else if arr[now] == 2 {
				singulars--
			}
		}

		score += singulars / 2
		if singulars%2 == 1 {
			score++
		}
		score += totalColors - (singulars / 2)
		fmt.Println(score)

	}

}
