package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ninja level 6 Exercises 8:")
	f := funkey()
	y := funkeyer()

	f(42)
	fmt.Println(y(43))
}

func funkey() func(i int) {
	return func(i int) {
		fmt.Println("I am a funkey function number", i)
	}
}

// This returns a function that takes in one arg and returns another.
func funkeyer() func(i int) string {
	return func(i int) string {
		return fmt.Sprintln("I am so funkey:", i)
	}
}
