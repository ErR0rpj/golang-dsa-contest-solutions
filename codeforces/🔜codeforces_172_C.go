package codeforces

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Codeforces_172_C() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Read the number of test cases
	tStr, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.TrimSpace(tStr))

	// Process each test case
	for tc := 0; tc < t; tc++ {
		// Read n and k
		firstLine, _ := reader.ReadString('\n')
		firstLineSplit := strings.Fields(firstLine)
		n, _ := strconv.Atoi(firstLineSplit[0])
		k, _ := strconv.Atoi(firstLineSplit[1])

		// Read the fish string
		fishStr, _ := reader.ReadString('\n')
		fishStr = strings.TrimSpace(fishStr)

		// Calculate the number of fishes Bob and Alice caught
		bobCount := 0
		aliceCount := 0
		for i := 0; i < n; i++ {
			if fishStr[i] == '1' {
				bobCount++
			} else {
				aliceCount++
			}
		}

		// If Bob cannot exceed Alice's score by k even if all his fishes are in the largest groups, return -1
		maxBobScore := (bobCount - 1) * (bobCount) / 2
		minAliceScore := (aliceCount - 1) * (aliceCount) / 2
		if maxBobScore-minAliceScore < k {
			fmt.Fprintln(writer, -1)
			continue
		}

		// Binary search for the minimum number of groups
		left, right := 1, n
		answer := n

		for left <= right {
			mid := (left + right) / 2
			groupSize := n / mid
			extra := n % mid

			bobScore := 0
			aliceScore := 0
			index := 0

			// Assign groups optimally
			for i := 0; i < mid; i++ {
				currentSize := groupSize
				if i < extra {
					currentSize++
				}

				// Calculate Bob's and Alice's scores for this group
				bobInGroup := 0
				aliceInGroup := 0
				for j := 0; j < currentSize; j++ {
					if fishStr[index] == '1' {
						bobInGroup++
					} else {
						aliceInGroup++
					}
					index++
				}

				bobScore += i * bobInGroup
				aliceScore += i * aliceInGroup
			}

			// Check if Bob's score exceeds Alice's score by at least k
			if bobScore-aliceScore >= k {
				answer = mid
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		// Output the answer
		fmt.Fprintln(writer, answer)
	}
}
