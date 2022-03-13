package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speek() {
	fmt.Println("Hi my name is", p.first, p.last, "and I am", p.age, "old")
}

func main() {
	fmt.Println("Ninja level 6 Exercises 3:")
	slim := person{
		first: "Slim",
		last:  "Shady",
		age:   34,
	}
	slim.speek()
}
