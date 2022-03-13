package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")

	a := 42
	fmt.Printf("The Value of a:\t\t%v\n", a)
	fmt.Printf("The pointer to a:\t%v\n", &a) // & referances the pointer "AND.. Whats the point..er?"
	fmt.Printf("The type of a:\t\t%T\n", a)
	fmt.Printf("The type of as pointer:\t%T\n\n", &a) // *int is a type!  v = *a is a defreance

	fmt.Println("Now lets pass the pointer around")
	var b *int // b is a pointer to int type.  its a pointer.. not an int.  It can point to an int.
	b = &a     // assign b to the pointer of a (they point to the same int or the same memeory register)
	fmt.Printf("The type of b:\t\t%T\n", b)
	fmt.Printf("The value of b:\t\t%v\n", b)
	fmt.Printf("The dereferanced value of b:\t%v\n\n", *b) // *b derefarances.  It gets whats in the memory the pointer points to

	fmt.Println("Now lets use the pointer with a function.")
	playWithTheSameMemory(&a) // We need to pass the pointer. So we use &
	fmt.Printf("The Value of a:\t\t%v\n", a)

}

// Accepts a pointer (type *int) to int
func playWithTheSameMemory(i *int) {
	fmt.Printf("The type of i:\t\t%T\n", i)
	fmt.Printf("The value of i:\t\t%v\n", i)
	v := *i // This sets v = to the value of i (dereferancing) at this point it makes a copy
	fmt.Printf("The pointer of v:\t%v --> Notice its a differnt memory address\n", &v)
	v += 10
	fmt.Printf("The value of v:\t\t%v\n", v)

	*i++ // Change the memory of the passed i (we can access and change that too.)
	// NOTE:  You should avoid doing this.  Passing references around should be done intentionally.
}
