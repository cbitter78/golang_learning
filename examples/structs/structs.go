package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

// Attach the function PrintDescription to the struct vehicle (a method)
func (v vehicle) PrintDescription() {
	fmt.Println("I am", v.color, "with", v.doors, "doors")
}

func main() {

	fmt.Println("Structs:")
	v := vehicle{
		doors: 2,
		color: "red",
	}

	tundra := truck{
		vehicle: vehicle{
			doors: 4,
			color: "green",
		},
		fourWheel: true,
	}

	cruse := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println("\t", v)
	fmt.Println("\tTundra", tundra)
	tundra.PrintDescription()
	fmt.Println("\tCurse:", cruse)
	cruse.PrintDescription()

}
