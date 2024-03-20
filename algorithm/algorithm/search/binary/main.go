package main

import "fmt"

func main() {

	arr := []int{1, 2, 5, 8, 11, 13}

	index := BinarySearch(arr, 11, 0, len(arr)-1)

	fmt.Println(index)
}
