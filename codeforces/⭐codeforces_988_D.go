// ⭐
package codeforces

import (
	"container/heap"
	"fmt"
	"main/data_structures"
	"main/templates"
)

// https://codeforces.com/contest/2037/problem/D
func Codeforces_988_D() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n, m, l int
		templates.TripleInput(&n, &m, &l)
		hurdles := templates.Inputnx2Array(n)
		powers := templates.Inputnx2Array(m)

		power := 1
		totalPowersTaken := 0
		currentHurdle := 0
		currentPower := 0

		var maxheap data_structures.MaxHeapInt
		heap.Init(&maxheap)

		for i := 1; i <= l; i++ {
			//power handle
			for ; currentPower < m && powers[currentPower][0] <= i; currentPower++ {
				heap.Push(&maxheap, powers[currentPower][1])
			}

			if currentHurdle < n && hurdles[currentHurdle][0] == i {
				hurdleSize := hurdles[currentHurdle][1] - hurdles[currentHurdle][0] + 1
				for power <= hurdleSize && maxheap.Len() > 0 {
					maxiPower := heap.Pop(&maxheap).(int)
					// fmt.Println("hurdlesize: ", hurdleSize)
					// fmt.Println("maxiPower: ", maxiPower)
					power += maxiPower
					totalPowersTaken += 1
				}

				if power <= hurdleSize {
					totalPowersTaken = -1
					break
				}
				i = hurdles[currentHurdle][1]
				currentHurdle++
			}
		}

		fmt.Println(totalPowersTaken)
	}
}
