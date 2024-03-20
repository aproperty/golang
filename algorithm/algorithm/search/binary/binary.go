package main

// 二分查找（前提必须在有序的数组里，查找 target）
// 如果找到 target，返回相应的索引 index
// 如果没有找到 target，返回-1

// 二分查找法-迭代法
func BinarySearchX(arr []int, target int, left int, right int) int {

	for left <= right {
		// mid := (right+left)/2
		// 下面的写法是为防止 right + left 溢出 int 最大值
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid

		} else if arr[mid] > target {
			right = mid - 1

		} else {
			left = mid + 1
		}
	}

	return -1
}

// 二分查找法-递归法
func BinarySearch(arr []int, target int, left int, right int) int {

	for left <= right {

		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid

		} else if arr[mid] > target {
			return BinarySearch(arr, target, left, mid-1)

		} else {
			return BinarySearch(arr, target, mid+1, right)
		}
	}

	return -1
}
