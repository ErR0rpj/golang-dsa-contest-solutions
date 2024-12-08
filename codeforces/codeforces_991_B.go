package codeforces

import "fmt"

func Codeforces_991_B() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)

		total := 0
		list := make([]int, n)
		for i := 0; i < n; i++ {
			var temp int
			fmt.Scan(&temp)
			list[i] = temp
			total += temp
		}

		if total%n != 0 {
			fmt.Println("No")
			continue
		}
		avg := total / n

		for i := 1; i < n-1; i++ {
			dif := 0
			if list[i-1] > avg {
				dif = list[i-1] - avg
				list[i-1] = list[i-1] - dif
				list[i+1] += dif
			} else {
				dif = avg - list[i-1]
				list[i-1] += dif
				list[i+1] = list[i+1] - dif
			}
		}

		ans := "Yes"
		for i := range list {
			if list[i] != avg {
				ans = "No"
				break
			}
		}

		fmt.Println(ans)

	}
}
