package sort

func QuickSort(arr []int64, start, end int) {

	if start < end {

		i, j := start, end

		key := arr[(start+end)/2]

		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}

			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			QuickSort(arr, start, j)
		}
		if i < end {
			QuickSort(arr, i, end)
		}
	} //

}
