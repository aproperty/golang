package mathutil

import (
	"math"
	"strconv"
	"strings"
)

var tenToAny map[int]string = map[int]string{
	0:  "0",
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "A",
	11: "B",
	12: "C",
	13: "D",
	14: "E",
	15: "F",
	16: "G",
	17: "H",
	18: "I",
	19: "J",
	20: "K",
	21: "L",
	22: "M",
	23: "N",
	24: "O",
	25: "P",
	26: "Q",
	27: "R",
	28: "S",
	29: "T",
	30: "U",
	31: "V",
}

// map 根据 value 找 key
func findKey(in string) int {
	result := -1
	for k, v := range tenToAny {
		if in == v {
			result = k
		}
	}
	return result
}

// 10 进制 转 任意进制，最多 32 进制
func DecimalToAny(num, n int) string {

	newNumStr := ""

	var remainder int
	var remainderStr string

	for num != 0 {
		remainder = num % n
		if 9 < remainder && remainder < 32 {
			remainderStr = tenToAny[remainder]
		} else {
			remainderStr = strconv.Itoa(remainder)
		}
		newNumStr = remainderStr + newNumStr
		num = num / n
	}

	return newNumStr
}

// 任意进制 转 10 进制
func AnyToDecimal(num string, n int) int {

	var newNum float64 = 0.0

	numSli := strings.Split(num, "")
	nNum := len(numSli) - 1

	for _, value := range numSli {

		tmp := float64(findKey(value))
		if tmp != -1 {
			newNum = newNum + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}

	return int(newNum)
}
