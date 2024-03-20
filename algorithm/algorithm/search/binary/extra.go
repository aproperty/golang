package main

// 在 循环递增数组 中查找某个值

func findX(arr []int, target int, left int, right int) int {

	// a := [8]int{4, 5, 6, 7, 8, 1, 2, 3}
	// search := 11
	// left := 0
	// right := len(a) - 1

	for left <= right {

		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		// if left == right {
		// 	return -1
		// }

		if arr[left] < arr[mid] { // 左侧为严格单调
			if arr[left] <= target && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else { // 右侧为严格单调
			if arr[mid] < target && target <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	} // end of `for`

	return -1
}
