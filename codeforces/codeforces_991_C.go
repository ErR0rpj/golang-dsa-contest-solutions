package codeforces

import "fmt"

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
			total += current % 9
			if current < 4 && current > 0 {
				list[current] += 1
			}
		}

		if total%9 == 0 {
			fmt.Println("Yes")
			continue
		}

		pend := 9 - (total % 9)
		if pend%2 == 1 {
			pend += 9
		}

		ans := "No"
		if pend == 2 && list[2] >= 1 {
			ans = "yes"
		} else if pend == 4 && list[2] >= 2 {
			ans = "yes"
		} else if pend == 6 && list[3] >= 1 || list[2] >= 3 {
			ans = "yes"
		} else if pend == 8 && ((list[2] >= 1 && list[3] >= 1) || (list[2] >= 4)) {
			ans = "yes"
		} else if pend == 10 && ((list[2] >= 2 && list[3] >= 1) || (list[2] >= 5)) {
			ans = "yes"
		} else if pend == 12 && ((list[2] >= 3 && list[3] >= 1) || (list[2] >= 6) || (list[3] >= 2)) {
			ans = "yes"
		} else if pend == 12 && ((list[2] >= 4 && list[3] >= 1) || (list[2] >= 7) || (list[3] >= 2 && list[2] >= 1)) {
			ans = "yes"
		} else if pend == 12 && ((list[2] >= 5 && list[3] >= 1) || (list[2] >= 8) || (list[3] >= 2 && list[2] >= 2)) {
			ans = "yes"
		}

		fmt.Println(ans)

	}
}
