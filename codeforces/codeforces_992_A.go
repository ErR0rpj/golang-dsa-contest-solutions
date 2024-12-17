package codeforces

import (
	"fmt"
	"main/templates"
	"math"
)

func Codeforces_992_A() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n, k int
		fmt.Scan(&n, &k)
		arr := make([]int, n)
		templates.InputArray(arr, n)

		ans := "No"
		val := -1
		for i := range arr {
			breaker := false
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				cur := int(math.Abs(float64(arr[i] - arr[j])))
				if cur%k == 0 {
					breaker = true
					break
				}
			}

			if breaker {
				continue
			} else {
				val = i + 1
				ans = "Yes"
				break
			}
		}

		fmt.Println(ans)
		if val != -1 {
			fmt.Println(val)
		}
	}
}
