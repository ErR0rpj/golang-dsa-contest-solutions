package main

import "fmt"

// https://codeforces.com/contest/2044/problem/E
func main() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var k, l1, r1, l2, r2 int64
		fmt.Scan(&k, &l1, &r1, &l2, &r2)

		var currentDivider int64
		var ans int64
		var start, end int64
		start, end = max(l1, l2), min(r1, r2)
		currentDivider = 1
		for l1 <= r2 {
			if start <= end {
				ans += (end / currentDivider) - (start / currentDivider)

				if start%currentDivider == 0 {
					ans += 1
				}
			}
			l1 = l1 * k
			r1 = r1 * k
			start, end = max(l1, l2), min(r1, r2)

			currentDivider = currentDivider * k
		}

		fmt.Println(ans)
	}
}
