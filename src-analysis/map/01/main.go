package main

import (
	"fmt"
	"sort"
)

func main() {

	myMap := make(map[int]string)
	myMap[1] = "红孩儿"
	myMap[2] = "牛魔王"
	myMap[3] = "白骨精"
	myMap[4] = "小钻风"
	myMap[5] = "黄袍怪"
	myMap[6] = "孔雀大明王"
	myMap[7] = "白毛鼠"

	// 获取所有的 key，取值后存储到切片
	keys := make([]int, 0, len(myMap))
	for k := range myMap {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	// 对 key 值进行排序
	// 内置函数 sort 包下的排序方法
	sort.Ints(keys)
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Println(key, "-->", myMap[key])
	}

}
