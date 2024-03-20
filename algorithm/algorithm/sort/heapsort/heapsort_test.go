package heapsort

import "testing"

func TestHeapSort(t *testing.T) {

	myArray := []int{89, 23, 4, 378, 23, 283, 56, 62, 8, 222, 64}

	result := HeapSort(myArray)

	t.Log(result)
}
