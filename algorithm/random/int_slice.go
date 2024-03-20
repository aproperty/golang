package randomutil

import (
	"math/rand"
	"time"
)

// 生成 int 数组
func GenerateRandomNumber(start int, end int, count int) []int {

	if end < start || (end-start) < count {
		return nil
	}

	nums := make([]int, 0) // 存放结果的 slice

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for len(nums) < count {
		num := r.Intn((end - start))
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
