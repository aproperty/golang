package main

// 使用二分法加速求解前导零数目
func bitsLen(x uint) int {

	clz := 0

	if x>>16 == 0 {
		clz += 16
		x <<= 16
	}

	if x>>24 == 0 {
		clz += 8
		x <<= 8
	}

	if x>>28 == 0 {
		clz += 4
		x <<= 4
	}

	if x>>30 == 0 {
		clz += 2
		x <<= 2
	}

	if x>>31 == 0 {
		clz++
	}

	return 32 - clz
}

// 使用分治法来加速求解二进制数位 1 的个数
func onesCount(num uint) int {
	// 0x55555555 等于 二进制的 0101 0101 0101 0101 0101 0101 0101 0101
	num = num&0x55555555 + num>>1&0x55555555

	// 0x33333333 等于 二进制的 0011 0011 0011 0011 0011 0011 0011 0011
	num = num&0x33333333 + num>>2&0x33333333

	// 0x0F0F0F0F 等于 二进制的 0000 1111 0000 1111 0000 1111 0000 1111
	num = num&0x0F0F0F0F + num>>4&0x0F0F0F0F

	// 0x00FF00FF 等于 二进制的 0000 0000 1111 1111 0000 0000 1111 1111
	num = num&0x00FF00FF + num>>8&0x00FF00FF

	// 0x0000FFFF 等于 二进制的 0000 0000 0000 0000 1111 1111 1111 1111
	num = num&0x0000FFFF + num>>16&0x0000FFFF

	return int(num)
}

func numberOfSteps(num int) int {

	if num == 0 {
		return 0
	}

	return bitsLen(uint(num)) - 1 +
		onesCount(uint(num))
}
