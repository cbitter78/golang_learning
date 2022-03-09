package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Ninja level 4 Exercises:")
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

func number_six() {

	fmt.Println("6: Using make to slice up all the states:")
	s := make([]string, 0, 50)
	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `Kentucky`, `Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	// Because i am lazy I am iterating over the source states to append to my slice
	for i := 0; i < len(states); i++ {
		s = append(s, strings.TrimSpace(states[i]))
		println(i, s, s[i])
	}
	fmt.Printf("\tlen:  %v\n", len(s))
	fmt.Printf("\tcap:  %v\n", len(s))
	fmt.Printf("\tdata: %v\n", s)

	for i := 0; i < len(s); i++ {
		fmt.Println("\t", i, s[i])
	}
}

func number_seven() {
	fmt.Println("7: Multi Dimensional Slice:")
	d := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	fmt.Println("\t", d)

	for row_i, row_v := range d {
		fmt.Println("\tRow: ", row_i)
		for i, v := range row_v {
			fmt.Println("\t\t", i, v)
		}
	}
}

func number_eight() {
	fmt.Println("8: Map of string Slices:")
	m := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Cars`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	for row_k, row_v := range m {
		fmt.Println("\t ", row_k)
		for i, v := range row_v {
			fmt.Println("\t\t", i, v)
		}
	}
}

func number_nine() {
	fmt.Println("9: Map of string Slices and add a record:")
	m := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Cars`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["me"] = []string{"fun", "games", "dogs"}
	for row_k, row_v := range m {
		fmt.Println("\t ", row_k)
		for i, v := range row_v {
			fmt.Println("\t\t", i, v)
		}
	}
}

func number_ten() {
	fmt.Println("10: Map of string Slices and delete a record:")
	m := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Cars`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["me"] = []string{"fun", "games", "dogs"}
	delete(m, "bond_james")
	for row_k, row_v := range m {
		fmt.Println("\t ", row_k)
		for i, v := range row_v {
			fmt.Println("\t\t", i, v)
		}
	}
}
