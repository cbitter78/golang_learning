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
	fmt.Println(y(43))
	fmt.Println(y(43))
	fmt.Println(y(43))
	fmt.Println(y(43))
	fmt.Println(y(43))
	fmt.Println(y(43))
	b := funkeyer()
	fmt.Println("--Reset the Closure --")
	fmt.Println(b(43))
	fmt.Println(b(43))
	fmt.Println("--Reset the Closure Even Only with Call back --")
	c := funkeyer()
	funkeyCallBackOnEven(c)
}

func funkey() func(i int) {
	return func(i int) {
		fmt.Println("I am a funkey function number", i)
	}
}

// This returns a function that takes in one arg and returns another.
func funkeyer() func(i int) string {
	y := 0
	return func(i int) string {
		y++
		return fmt.Sprintln("I am so funkey:", i, y)
	}
}

func funkeyCallBackOnEven(x func(i int) string) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even Number:", i, x(43))
		}
	}
}
