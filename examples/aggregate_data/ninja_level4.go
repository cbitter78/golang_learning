package main

import (
	"fmt"
)

func main() {

	fmt.Println("Ninja level 4 Exercises:")
	number_one()
	number_two()
	number_three()
	number_four()
	number_five()
}

func number_one() {
	fmt.Println("1: Create an array and print its index and values then the array type:")
	a := [5]int{1, 3, 5, 8, 13}
	for i, v := range a {
		fmt.Println("\t", i, v)
	}
	fmt.Printf("\t%T\n", a)
}

func number_two() {
	fmt.Println("2: Create an slice and print its index and values then the array type:")
	s := []int{1, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987}
	for i, v := range s {
		fmt.Println("\t", i, v)
	}
	fmt.Printf("\t%T\n", s)
}

func number_three() {
	fmt.Println("3: Slice the slice:")
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("\tFull Slice:  ", s)
	fmt.Println("\tFirst Five:  ", s[:5])
	fmt.Println("\tFirst Five:                 ", s[5:])
	fmt.Println("\tFirst Mid 1:       ", s[2:7])
	fmt.Println("\tFirst Mid 2:    ", s[1:6])
}

func number_four() {
	fmt.Println("4: Append to the slice:")
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("\t", s)
	s = append(s, 52)
	fmt.Println("\t", s)
	s = append(s, 53, 54, 55)
	fmt.Println("\t", s)
	y := []int{56, 57, 58, 59, 60}
	s = append(s, y...)
	fmt.Println("\t", s)
}

func number_five() {
	fmt.Println("5: Delete From the slice:")
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("\t", s)
	s = append(s[:3], s[6:]...)
	fmt.Println("\t", s)
}
