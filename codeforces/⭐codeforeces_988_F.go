package codeforces

import (
	"fmt"
	"main/templates"
)

// https://codeforces.com/contest/2037/problem/F
func Codeforces_988_F() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n, m, k int
		templates.TripleInput(&n, &m, &k)

		health := make([]int, n)
		templates.InputArray(health, n)
		position := make([]int, n)
		templates.InputArray(position, n)

		divider := make([]int, (m*2)-1)
		counter := 1
		for i := range divider {
			divider[i] = counter
			if i >= m-1 {
				counter--
			} else {
				counter++
			}
		}

		ans := -1
		miniAttack, maxiAttack, currentAttack := 1, 1, 1
		for i := range health {
			if health[i] > maxiAttack {
				maxiAttack = health[i]
			}
		}

		for miniAttack <= maxiAttack {
			currentAttack = (miniAttack + maxiAttack) / 2

		}

	}
}
