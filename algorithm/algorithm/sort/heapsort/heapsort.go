package heapsort

func createHeap(arr []int, end int) {

	for start := end / 2; start >= 0; start-- {
		maxHeap(start, end, arr)
	}
}

func HeapSort(c []int) []int {

	n := len(c) - 1

	createHeap(c, n)

	for i := n; i > 0; i-- {

		if c[0] > c[i] {
			c[0], c[i] = c[i], c[0]

			maxHeap(0, i-1, c)
		}
	}

	return c
}
