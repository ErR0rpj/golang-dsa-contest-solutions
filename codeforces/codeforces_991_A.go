package codeforces

import "fmt"

func Codeforces_991_A() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n, m int
		fmt.Scan(&n, &m)

		list := make([]string, n)
		for i := 0; i < n; i++ {
			var temp string
			fmt.Scan(&temp)
			list[i] = temp
		}

		ans := 0
		currentLen := 0
		for i := range list {
			if currentLen+len(list[i]) <= m {
				ans += 1
				currentLen += len(list[i])
			} else {
				break
			}
		}

		fmt.Println(ans)
	}
}
