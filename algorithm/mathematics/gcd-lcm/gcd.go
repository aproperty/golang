package main

// 最大公约数（greatest common divisor）

// // 穷举法：最大公约数
// func gcdNormal(x, y int) int {

// 	var n int
// 	if x > y {
// 		n = y
// 	} else {
// 		n = x
// 	}

// 	for i := n; i >= 1; i-- {
// 		if x%i == 0 && y%i == 0 {
// 			return i
// 		}
// 	}

// 	return 1
// }

// 欧几里得 辗转相除法：
// gcd(x,y) 表示 x 和 y 的最大公约数。进入运算时：x != 0, y != 0，
// 本质上就是：不断转换成求等价更小数的最大公约数。
// 如果 x%y = 0，返回 y，即最大公约数。
// gcd(x,y) = gcd(y,x%y)

// 证明：设 k = x/y，b = x%y  则：x = ky+b
// 如果 n 能够同时整除 x 和 y，则 (y%n) = 0，(ky+b)%n = 0，则 b%n = 0，即 n 也同时能够整除 y 和 b。
// 由上得出：同时能够整除 y 和 (b=x%y) 的数，也必然能够同时整除 x 和 y。
// 故而 gcd(x,y) = gcd(y,x%y)。
// 当 (b = x%y) = 0，即 y 可以整除 x，这时的 y 也就是所求的最大公约数了。

// 附上两条重要性质：gcd(a,b) = gcd(b,a)， gcd(-a,b)=gcd(a,b)

// 递归写法，进入运算是 x 和 y 都不为 0
func gcd(x, y int) int {

	tmp := x % y

	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}

// 非递归写法
func gcdx(x, y int) int {

	var tmp int

	for {

		tmp = (x % y)

		if tmp > 0 {
			x = y
			y = tmp

		} else {
			return y
		}
	}
}
