package codeforces

import (
	"fmt"
	"main/templates"
)

// func ableToKill(position []int, health []int, divider []int, m int, k int, startInterval int, endInterval int) bool {
// 	// dividerPos := 0
// 	// for i := startInterval ; i<len(divider); i++ {
// 	// 	if startInterval <
// 	// }

// 	return false
// }

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
		// 5 5 3
		// 7 7 7 7 7
		// 1 2 3 4 5
		for miniAttack <= maxiAttack {
			var intervals [][]int
			currentAttack = (miniAttack + maxiAttack) / 2
			for i := range health {
				powerReq := health[1] / currentAttack
				if health[i]%currentAttack > 0 {
					powerReq += 1
				}
				if powerReq <= m {
					endInterval := position[i] + (m - powerReq) - m
					startInterval := position[i] - (m - powerReq) - m
					interval := make([]int, 2)
					interval[0] = startInterval
					interval[1] = endInterval
					intervals = append(intervals, interval)
				}
			}
			possible := false
			for i := range intervals {
				if i == 0 {
					possible = true
					break
				}
			}
			if possible {
				maxiAttack = currentAttack - 1
				if ans == -1 {
					ans = currentAttack
				}
				if currentAttack < ans {
					ans = currentAttack
				}
			} else {
				miniAttack = currentAttack + 1
			}
		}
		fmt.Println(ans)
	}
}
