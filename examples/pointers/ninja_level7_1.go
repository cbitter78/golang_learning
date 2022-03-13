package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ninja level 7 Exercises 1:")
	v := "Go Ninjas Go..."
	fmt.Printf("The address is:\t%v\n", &v)
	fmt.Printf("The type is:\t%T\n", &v)
	fmt.Printf("The value is:\t`%v` and it is of type `%T`\n", v, v)
}
