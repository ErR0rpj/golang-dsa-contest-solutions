package codeforces

import "fmt"

//https://codeforces.com/contest/2050/problem/C
// Always gives wrong ans for test case 3, 451st value.
func Codeforces_991_C() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)

		total := 0
		list := make([]int, 4)
		for n > 0 {
			current := n % 10
			n = n / 10
			total += current
			if current < 4 && current > 0 {
				list[current] += 1
			}
		}

		ans := "No"
		for i := 0; i <= min(9, list[2]); i++ {
			for j := 0; j <= min(9, list[3]); j++ {
				if (total+(2*i)+(6*j))%9 == 0 {
					ans = "Yes"
					break
				}
			}
			if ans == "Yes" {
				break
			}
		}

		fmt.Println(ans)

	}
}
