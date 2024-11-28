package main

import (
	"fmt"
)

// Takes the input for predefined length n array
func InputArray(input []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
	}
}

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
func main() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)

		ans := "YES"
		arr := make([]int, n)
		InputArray(arr, n)

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
