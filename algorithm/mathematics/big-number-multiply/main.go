package main

import "fmt"

// BigMulti 大数相乘
func BigMulti(a, b string) string {

	if a == "0" || b == "0" {
		return "0"
	}

	// string 转换成 []byte，容易取得相应位上的具体值
	bsi := []byte(a)
	bsj := []byte(b)

	temp := make([]int, len(bsi)+len(bsj))

	// 两数相乘，结果位数不会超过 两乘数 位数和，
	// 即 temp 的长度只可能为 len(num1)+len(num2) 或 len(num1)+len(num2)-1
	// 选最大的，免得位数不够

	for i := 0; i < len(bsi); i++ {
		for j := 0; j < len(bsj); j++ {
			temp[i+j+1] += int(bsi[i]-'0') * int(bsj[j]-'0')
		}
	}

	// 统一处理进位
	for k := len(temp) - 1; k > 0; k-- {
		temp[k-1] += temp[k] / 10 // 对该结果进位（进到前一位）
		temp[k] = temp[k] % 10    // 对个位数保留
	}

	if temp[0] == 0 {
		temp = temp[1:]
	}

	// 将 []int 类型的 temp 转成 []byte 类型,
	// 因为在未处理进位的情况下，temp 每位的结果可能超过 255(go中，byte类型实为uint8，最大为255),所以temp选用[]int类型
	// 但在处理完进位后，不再会出现溢出
	res := make([]byte, len(temp)) // res 存放最终结果的 ASCII 码

	for i := 0; i < len(temp); i++ {
		res[i] = byte(temp[i] + '0')
	}
	return string(res)
}

func main() {

	a := "999999999999"
	b := "111111111111"

	c := BigMulti(a, b)

	fmt.Println(c)
}
