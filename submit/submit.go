package main

import (
	"fmt"
	"math"
)

func InputArray(input []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
	}
}

func main() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n, k int
		fmt.Scan(&n, &k)
		arr := make([]int, n)
		InputArray(arr, n)

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
