package graph

import (
	"fmt"
	"sort"
)

// https://www.hackerearth.com/practice/algorithms/graphs/graph-representation/practice-problems/algorithm/split-the-bill-3-5a0690ff/
func SplitBill() {
	var n, m int

	fmt.Scan(&n, &m)
	person := make([]int, n+1)

	for ; m > 0; m-- {
		var id string
		fmt.Scan(&id)

		var payerCount, splitCount int
		fmt.Scan(&payerCount, &splitCount)

		for i := 0; i < payerCount; i++ {
			var payer, amount int
			fmt.Scan(&payer, &amount)

			person[payer] += amount
		}

		for i := 0; i < splitCount; i++ {
			var payer, amount int
			fmt.Scan(&payer, &amount)

			person[payer] -= amount
		}
	}

	var receiver, payer [][]int
	for i := 1; i <= n; i++ {
		if person[i] > 0 {
			temp := make([]int, 2)
			temp[0] = i
			temp[1] = person[i]
			receiver = append(receiver, temp)
		} else if person[i] < 0 {
			temp := make([]int, 2)
			temp[0] = i
			temp[1] = person[i]
			payer = append(payer, temp)
		}
	}

	sort.Slice(receiver, func(i, j int) bool {
		return receiver[i][0] < receiver[j][0]
	})
	sort.Slice(payer, func(i, j int) bool {
		return payer[i][0] < payer[j][0]
	})

	var ans [][]int
	currentReceiverPos := 0
	for i := 0; i < len(payer); i++ {
		payment := payer[i][1] * -1

		for payment > 0 {
			receiverment := receiver[currentReceiverPos][1]
			temp := make([]int, 3)
			temp[0] = payer[i][0]
			temp[1] = receiver[currentReceiverPos][0]
			if receiverment <= payment {
				temp[2] = receiverment
				payment = payment - receiverment
				receiver[currentReceiverPos][1] = 0
				currentReceiverPos++
			} else {
				temp[2] = payment
				receiver[currentReceiverPos][1] -= payment
				payment = 0
			}
			ans = append(ans, temp)
		}
	}

	for i := 0; i < len(ans); i++ {
		fmt.Print(ans[i][0], " ", ans[i][1], " ", ans[i][2])
		fmt.Println("")
	}
}
