package templates

import "fmt"

// Takes the input for predefined length n array
func InputArray(input []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
	}
}

// Takes the input for nx2 not pre maked array
func Inputnx2Array(n int) [][]int {
	var input [][]int
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a)
		fmt.Scan(&b)

		arr := []int{a, b}
		input = append(input, arr)
	}
	return input
}
