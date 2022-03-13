package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ninja level 6 Exercises 1:")
	x := foo()
	y, z := bar()
	fmt.Println(x)
	fmt.Println(y, z)
}

func foo() string {
	return "The Foo Function"
}

func bar() (string, int) {
	return "The Bar Function", 1075
}
