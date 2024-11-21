package templates

import "fmt"

func InputArray(input []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
	}
}
