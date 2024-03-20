package main

import (
	"fmt"
)

func main() {

	x, y := 12, 15

	// fmt.Println("gcdNormal=", gcdNormal(x, y))

	fmt.Println("gcd=", gcd(x, y))
	fmt.Println("gcdx=", gcdx(x, y))

	// fmt.Println("lcmNormal=", lcmNormal(x, y))
	fmt.Println("lcm=", lcm(x, y))
}
