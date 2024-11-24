package data_structures

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
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeapInt) Push(x any) {
	*h = append(*h, x.(int))
}

func (h MaxHeapInt) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
