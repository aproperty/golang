package sort

func BubbleSort(values []int) {

	n := len(values)

	for i := 0; i < n-1; i++ {

		for j := i + 1; j < n; j++ {

			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}

		}

	}

}
