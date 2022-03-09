package main

import (
	"fmt"
)

func main() {

	fmt.Println("Ninja level 5 Exercises:")
	number_one()
	number_two()
	number_three()
	number_four()
	number_five()
	number_six()
	number_seven()
	number_eight()
	number_nine()
	number_ten()
}

func number_one() {
	fmt.Println("1: Create struct and print it")
	type person struct {
		first_name string
		last_name  string
		fav_flavor []string
	}

	p1 := person{
		first_name: "Bob",
		last_name:  "Jones",
		fav_flavor: []string{"chocolate", "cheescake", "rockyroad"},
	}

	p2 := person{
		first_name: "crazy",
		last_name:  "willy",
		fav_flavor: []string{"Free Harassment"}, // This is an ice cream flavor made by Crazy Willy in Front Royal VA.  :)
	}

	for i, p := range []person{p1, p2} {
		fmt.Printf("\t%v: %v %v likes these flavors of ice cream:\n", i, p.last_name, p.last_name)
		for _, f := range p.fav_flavor {
			fmt.Println("\t\t", f)
		}
	}

}

func number_two() {
	fmt.Println("1: Create struct, add the values to a map and print with range")
	type person struct {
		first_name string
		last_name  string
		fav_flavor []string
	}
	people := map[string]person{
		"Jones": {
			first_name: "Bob",
			last_name:  "Jones",
			fav_flavor: []string{"chocolate", "cheescake", "rockyroad"},
		},
		"Willy": {
			first_name: "crazy",
			last_name:  "Willy",
			fav_flavor: []string{"Free Harassment"},
		},
	}

	// Add in a record to the map for extra credit.  :)
	p := person{
		first_name: "Sam",
		last_name:  "Nickels",
		fav_flavor: []string{"Cookie Dough"},
	}
	people[p.last_name] = p

	for last_name, p := range people {
		fmt.Printf("\t%v: %v %v likes these flavors of ice cream:\n", last_name, p.last_name, p.last_name)
		for _, f := range p.fav_flavor {
			fmt.Println("\t\t", f)
		}
	}

}

func number_three() {
	fmt.Println("3: Embeded type")
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
	fmt.Println("\t", tundra)
	fmt.Println("\t", cruse)
}

func number_four() {

}

func number_five() {

}

func number_six() {

}

func number_seven() {

}

func number_eight() {

}

func number_nine() {

}

func number_ten() {

}
