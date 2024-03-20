package main

// 最小公倍数（least common multiple）

// // 穷举写法：最小公倍数
// func lcmNormal(x, y int) int {

// 	var i = x
// 	if x < y {
// 		i = y
// 	}

// 	var top int = x * y
// 	for ; i <= top; i++ {
// 		if i%x == 0 && i%y == 0 {
// 			return i
// 		}
// 	}

// 	return top
// }

// 公式解法：
// 最小公倍数 = 两数之积/最大公约数
func lcm(x, y int) int {
	return x * y / gcd(x, y)
}
