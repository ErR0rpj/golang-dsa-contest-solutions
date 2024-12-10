package codeforces

import (
	"fmt"
	"sort"
)

// https://codeforces.com/contest/2042/problem/C
func Codeforces_172_C() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n, k int
		fmt.Scan(&n, &k)
		var s string
		fmt.Scan(&s)

		prefSum := make([]int, n)
		total := 0
		for i := n - 1; i >= 1; i-- {
			if s[i] == '1' {
				total += 1
			} else {
				total -= 1
			}
			prefSum[i] = total
		}

		sort.SliceStable(prefSum, func(i, j int) bool {
			return prefSum[i] > prefSum[j]
		})

		currentSum := 0
		ans := 1
		for i := 0; i < n; i++ {
			if currentSum >= k || prefSum[i] <= 0 {
				break
			} else {
				currentSum += prefSum[i]
				ans += 1
			}
		}

		if currentSum < k {
			ans = -1
		}

		fmt.Println(ans)

	}
}
