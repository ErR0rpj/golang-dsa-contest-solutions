package codeforces

import "fmt"

func Codeforces_993_C() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var m, a, b, c int
		fmt.Scan(&m, &a, &b, &c)
		row1 := m
		row2 := m
		ans := 0

		row1 = row1 - min(m, a)
		ans += min(m, a)

		row2 = row2 - min(m, b)
		ans += min(m, b)

		ans += min(row1+row2, c)

		fmt.Println(ans)
	}
}
