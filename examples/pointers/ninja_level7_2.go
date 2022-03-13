package main

import (
	"fmt"
	"strings"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) printPerson() {
	fmt.Println("\tfirst name:\t", p.first)
	fmt.Println("\tfirst name:\t", p.last)
	fmt.Println("\tage:\t\t", p.age)
}

func main() {
	fmt.Println("Ninja level 7 Exercises 2:")
	p := person{
		first: "jeff",
		last:  "foxworthy",
		age:   45,
	}
	p.printPerson()
	changeMe(&p)
	fmt.Println("\n  Change person using pointers..")
	p.printPerson()

}

func changeMe(p *person) {
	(*p).first = strings.Title(strings.ToLower((*p).first)) // Dereferancing all of p
	p.last = strings.Title(strings.ToLower((*p).last))      // Accessing last on p (you can do this too.)
	p.age++
}
