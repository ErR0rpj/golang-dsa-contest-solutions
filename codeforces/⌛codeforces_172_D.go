package codeforces

import "fmt"

func predictorFinder(arr [][]int, index int) []int {
	var predictors []int
	start := arr[index][0]
	end := arr[index][1]

	for i := 0; i < len(arr); i++ {
		if i == index {
			continue
		} else {
			if arr[i][0] <= start && arr[i][1] >= end {
				predictors = append(predictors, i)
				// fmt.Println("predictor for", index, "is", i)
			}
		}
	}

	return predictors
}

func Codeforces_172_D() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)

		arr := make([][]int, n)
		for i := 0; i < n; i++ {
			temp := make([]int, 2)
			fmt.Scan(&temp[0], &temp[1])
			arr[i] = temp
		}

		predictors := make([][]int, n)
		for i := 0; i < n; i++ {
			temp := predictorFinder(arr, i)
			predictors[i] = temp
		}

		for i := 0; i < n; i++ {
			if len(predictors[i]) == 0 {
				fmt.Println(0)
			} else {
				predictor := predictors[i]
				start := arr[predictor[0]][0]
				end := arr[predictor[0]][1]
				ans := 0

				for j := 0; j < len(predictor); j++ {
					start = max(start, arr[predictor[j]][0])
					end = min(end, arr[predictor[j]][1])

					if start > end {
						break
					}
				}

				if start > end {
					fmt.Println(ans)
				} else {
					ans = end - start - arr[i][1] + arr[i][0]
					fmt.Println(ans)
				}
			}
		}

	}
}
