package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ninja level 6 Exercises 2:")
	xi := []int{1, 3, 5, 7, 9}
	x := foo2(xi...)
	y := bar2(xi)
	fmt.Println(x)
	fmt.Println(y)
}

func foo2(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar2(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
