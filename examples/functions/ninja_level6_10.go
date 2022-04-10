package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ninja level 6 Exercises 10:")

	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println("fibonacci #", i, ":\t", fib())
	}

	fib2 := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println("fibonacci #", i, ":\t", fib2())
	}

	c := compound()
	for i := 0; i < 100; i++ {
		fmt.Println("compound #", i, ":\t", c(i))
	}

	// Anonymous closuer
	for i := 0; i < 7; i++ {
		v := 0
		func(ii int) {
			v += ii
			fmt.Println("anonymous #", i, ":\t", v)
		}(i)
	}
}

// Closure function that accepts an argument
func compound() func(i int) int {
	v := 0
	return func(i int) int {
		v = (v + i)
		return v
	}
}

// Closure function with no arguments
func fibonacci() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}
}
