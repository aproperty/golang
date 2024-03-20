package main

import "fmt"

// func runningSum(nums []int) []int {

// 	sums := make([]int, len(nums))
// 	copy(sums, nums)

// 	for i := 1; i < len(nums); i++ {
// 		sums[i] += sums[i-1]
// 	}
// 	return sums
// }

func runningSum(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

func main() {

	one := []int{1, 2, 3, 4}
	oneRes := runningSum(one)
	fmt.Println(oneRes) // 1,3,6,10

	two := []int{1, 1, 1, 1, 1}
	twoRes := runningSum(two)
	fmt.Println(twoRes) // 1,2,3,4,5

	three := []int{3, 1, 2, 10, 1}
	threeRes := runningSum(three)
	fmt.Println(threeRes) // 3,4,6,16,17
}
