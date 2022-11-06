package main

import "fmt"

type Human struct {
	Name string
}

var people = Human{Name: "Zhangsan"}

func main() {

	fmt.Printf("%v \n", people)
	fmt.Printf("%+v \n", people)
	fmt.Printf("%#v \n", people)
	fmt.Printf("%T \n", people)
}
