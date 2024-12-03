package codeforces

import (
	"fmt"
	"main/templates"
	"slices"
)

func Codeforces_172_A() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n, k int
		fmt.Scan(&n, &k)

		arr := make([]int, n)
		templates.InputArray(arr, n)
		slices.Sort(arr)

		currentSum := 0
		for i := n - 1; i >= 0; i-- {
			currentSum += arr[i]
			if currentSum > k {
				currentSum -= arr[i]
				break
			} else if currentSum < k {
				continue
			} else {
				break
			}
		}

		ans := k - currentSum
		fmt.Println(ans)
	}
}
