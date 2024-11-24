package main

import (
	"container/heap"
	"fmt"
)

func TripleInput(a *int, b *int, c *int) {
	fmt.Scan(a)
	fmt.Scan(b)
	fmt.Scan(c)
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

type MaxHeapInt []int

func (h MaxHeapInt) Len() int {
	return len(h)
}

func (h MaxHeapInt) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *MaxHeapInt) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *MaxHeapInt) Push(x any) {
	*h = append(*h, x.(int))
}

func (h MaxHeapInt) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// https://codeforces.com/contest/2037/problem/D
// This question didnt got accepted with go and instead same logic worked with c++
func main() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n, m, l int
		TripleInput(&n, &m, &l)
		hurdles := Inputnx2Array(n)
		powers := Inputnx2Array(m)

		power := 1
		totalPowersTaken := 0
		currentHurdle := 0
		currentPower := 0

		var maxheap MaxHeapInt
		heap.Init(&maxheap)

		for currentHurdle < n {
			//power handle
			for ; currentPower < m && powers[currentPower][0] < hurdles[currentHurdle][0]; currentPower++ {
				heap.Push(&maxheap, powers[currentPower][1])
			}

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
			currentHurdle++

		}

		fmt.Println(totalPowersTaken)
	}
}
