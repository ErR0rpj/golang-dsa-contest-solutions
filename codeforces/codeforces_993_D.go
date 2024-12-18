package codeforces

import (
	"fmt"
	"main/templates"
)

func Codeforces_993_D() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)
		arr := make([]int, n)
		templates.InputArray(arr, n)
		temp := make(map[int]int)
		var leftovers []int

		for _, val := range arr {
			temp[val]++
		}
		for i := 1; i <= n; i++ {
			if temp[i] == 0 {
				leftovers = append(leftovers, i)
			}
		}

		j := 0
		ans := make([]int, n)
		for i := range arr {
			if temp[arr[i]] != 0 {
				ans[i] = arr[i]
				temp[arr[i]] = 0
			} else {
				ans[i] = leftovers[j]
				j++
			}
		}

		for _, val := range ans {
			fmt.Print(val, " ")
		}
		fmt.Println("")

	}
}
