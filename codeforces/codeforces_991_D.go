package codeforces

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func maximizeLexicographically(s string) string {
	// Convert string to byte slice for easier manipulation
	arr := []byte(s)
	n := len(arr)

	// We can't modify leftmost digit or zero digits
	for i := n - 1; i > 0; i-- {
		// Find the rightmost digit that can be decreased
		for j := i; j > 0; j-- {
			if arr[j] > '0' {
				// Decrease the digit
				arr[j]--

				// Swap with left digit
				arr[j], arr[j-1] = arr[j-1], arr[j]

				// Re-sort subarray from j to keep lexicographically maximum
				for k := j; k < n-1; k++ {
					if arr[k] < arr[k+1] {
						arr[k], arr[k+1] = arr[k+1], arr[k]
					} else {
						break
					}
				}

				break
			}
		}
	}

	return string(arr)
}

func Codeforces_991_D() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read number of test cases
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	// Process each test case
	for i := 0; i < t; i++ {
		scanner.Scan()
		s := scanner.Text()
		result := maximizeLexicographically(s)
		fmt.Println(result)
	}
}
