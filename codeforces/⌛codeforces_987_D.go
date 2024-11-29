package codeforces

import (
	"fmt"
	"main/templates"
)

func findHighest(arr []int, highest []int, current int) int {
	ans := current
	for current != highest[current] {
		current = highest[current]

	}
	highest[ans] = current
	return arr[current]
}

func Codeforces_987_D() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)
		arr := make([]int, n)
		templates.InputArray(arr, n)

		highest := make([]int, n)
		for i := 1; i < n; i++ {
			if arr[highest[i-1]] > arr[i] {
				highest[i] = highest[i-1]
			} else {
				highest[i] = i
			}
		}

		var tempHigh [][]int
		for i := n - 1; i >= 0; i-- {
			flag := true
			for j := range tempHigh {
				if tempHigh[j][0] > arr[highest[i]] {
					thekedar := tempHigh[j][1]
					if arr[thekedar] < arr[i] {
						highest[i] = highest[thekedar]
						flag = false
						break
					}
				} else if tempHigh[j][0] == arr[highest[i]] {
					thekedar := tempHigh[j][1]
					if arr[thekedar] > arr[i] {
						tempHigh[j][1] = i
					}
					flag = false
					break
				} else {
					break
				}
			}

			if flag {
				temp := make([]int, 2)
				temp[0] = arr[highest[i]]
				temp[1] = i
				tempHigh = append(tempHigh, temp)
			}
		}

		for i := range arr {
			fmt.Print(findHighest(arr, highest, i), " ")
		}
		fmt.Println("")
	}
}
