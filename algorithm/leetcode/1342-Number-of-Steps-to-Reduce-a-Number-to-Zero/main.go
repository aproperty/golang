package main

import "fmt"

func main() {

	sliceForTest := []int{14, 8, 123}

	for _, v := range sliceForTest {
		result := numberOfSteps(v)
		fmt.Println(result)
	}
}
