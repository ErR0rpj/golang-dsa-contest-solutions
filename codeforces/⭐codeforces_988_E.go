package codeforces

import "fmt"

func ask(pointer1 int, pointer2 int, steps *int) (sum int) {
	fmt.Println("?", pointer1, pointer2)
	fmt.Scan(&sum)
	*steps++

	return sum
}

//https://codeforces.com/contest/2037/problem/E
func Codeforces_988_E() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)

		ans := make([]int, n+1)
		pointer1 := 1
		pointer2 := n
		currentSum, previousSum := -1, -1
		steps := 0
		filled := 0
		currentSum = ask(pointer1, pointer2, &steps)
		pointer1++

		if currentSum == 0 {
			fmt.Println("! IMPOSSIBLE")
			continue
		}

		for currentSum != 0 && pointer2 > pointer1 {
			previousSum = currentSum
			currentSum = ask(pointer1, pointer2, &steps)

			if currentSum < previousSum {
				ans[pointer1-1] = 0
				filled += 1
				pointer1++
			} else if currentSum == previousSum {
				ans[pointer1-1] = 1
				filled += 1
				pointer1++
			}
		}

		if currentSum != 0 {
			pointer1--
			ans[pointer1] = 0
			ans[pointer2] = 1
			filled += 2
		} else {
			ans[pointer1-1] = 1
			filled += 1
			pointer1 = pointer1 - 2
			currentSum = previousSum
			pointer2--

			for currentSum != 0 && pointer2 > pointer1 {
				currentSum = ask(pointer1, pointer2, &steps)

				if currentSum < previousSum {
					ans[pointer2+1] = 1
					filled += 1
					previousSum = currentSum
					pointer2--
				} else if currentSum == previousSum {
					ans[pointer2+1] = 0
					filled += 1
					pointer2--
				}
			}
		}

		if filled < n {
			fmt.Println("! IMPOSSIBLE")
		} else {
			fmt.Print("! ")
			for i := 1; i <= n; i++ {
				fmt.Print(ans[i])
			}
			fmt.Println("")
		}

	}
}
