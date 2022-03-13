package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ninja level 6 Exercises 6:")
	func(i int) {
		fmt.Println(i)
	}(42)
}
