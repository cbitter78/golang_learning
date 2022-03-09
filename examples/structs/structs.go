package main

import (
	"fmt"
)

func main() {

	fmt.Println("Structs:")

	// Remember structs do not have commas between the fields!
	type vehicle struct {
		doors int
		color string
	}

}
