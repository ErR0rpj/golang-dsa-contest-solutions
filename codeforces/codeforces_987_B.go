package codeforces

import (
	"fmt"
	"main/templates"
)

func isSwapable(arr []int, index int) bool {
	currentNo := arr[index-1]
	dif := currentNo - index

	if dif != -1 && dif != 1 {
		return false
	}

	if arr[currentNo-1] == index {
		return true
	}
	return false

}

// https://codeforces.com/contest/2031/problem/B
func Codeforces_987_B() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)

		ans := "YES"
		arr := make([]int, n)
		templates.InputArray(arr, n)

		for i := range arr {
			index := i + 1
			if index == arr[i] {
				continue
			}
			flag := isSwapable(arr, index)
			if !flag {
				ans = "NO"
				break
			} else {
				ans = "YES"
			}
		}

		fmt.Println(ans)

	}
}
