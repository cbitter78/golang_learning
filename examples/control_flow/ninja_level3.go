package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ninja level 3 Exercises:")
	number_one()
	number_tow()
	number_three()
	number_four()
	number_five()
	number_six()
	number_seven(1, 2)
	number_seven(2, 2)
	number_seven(2, 1)
	number_eight()
	s := []string{"one", "two", "three", "four", "five", "hotdog"}
	for i := 0; i < 6; i++ {
		number_nine((s[i]))
	}
	number_ten()
}

func number_one() {
	fmt.Println("Print all numbers 1 through 10000:")
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func number_tow() {
	fmt.Println("Print all runes for all upper letters of the alphabet each three times:")
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for r := 0; r < 3; r++ {
			fmt.Printf("\t%#U\n", i)
		}

	}
}

func number_three() {
	fmt.Println("Print out all the years you have been alive if you were born in 1985:")
	year, _, _ := time.Now().Date()
	birth_year := 1985
	y := birth_year

	for y <= year {
		println(y)
		y++
	}
	fmt.Printf("You have been alive %v years\n", (y - birth_year))
}

func number_four() {
	fmt.Println("Print out all the years you have been alive if you were born in 1986 (Using for {}):")
	year, _, _ := time.Now().Date()
	birth_year := 1986
	y := birth_year

	for y <= year {
		println(y)
		if y == year {
			break
		}
		y++
	}
	fmt.Printf("You have been alive %v years\n", (y - birth_year))
}

func number_five() {
	fmt.Println("Print the remaindar of all numbers between 10 - 100 when devided by 4:")
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v\t / 4 remainder: %v\n", i, i%4)
	}
}

func number_six() {
	fmt.Println("Show the if statement in action:")
	x := 1
	y := 2

	if y == x {
		fmt.Printf("\tY is equal to x \t\t\t y:%v x:%v\n", y, x)
	} else {
		fmt.Printf("\tY is NOT equal to x \t\t\t y:%v x:%v\n", y, x)
	}
}

func number_seven(y int, x int) {
	fmt.Println("Show the if, elseif and else statement in action:")

	if y == x {
		fmt.Printf("\tY is equal to x \t\t\t y:%v x:%v\n", y, x)
	} else if y > x {
		fmt.Printf("\tY is greater than x \t\t\t y:%v x:%v\n", y, x)
	} else {
		fmt.Printf("\tY is NOT equal to or grater than x \t y:%v x:%v\n", y, x)
	}
}

func number_eight() {
	fmt.Printf("Use the Switch Statement with no switch expression:\n\t")
	switch {
	case true:
		fmt.Printf("Case TRUE\n")
	case false:
		fmt.Printf("Case FALSE\n")
	}
}

func number_nine(x string) {
	fmt.Printf("Use the Switch Statement with no switch expression:\n\t")
	switch x {
	case "one", "two":
		fmt.Printf("Case one or two  \t%v\n", x)
	case "three":
		fmt.Printf("Case three       \t%v\n", x)
	case "four", "five":
		fmt.Printf("Case four or five\t%v\n", x)
	default:
		fmt.Printf("Case default     \t%v\n", x)

	}
}

func number_ten() {
	fmt.Println("Show not, and, or:")
	fmt.Printf("\tture && true:  %v\t\n", (true && true))
	fmt.Printf("\tture && false: %v\t\n", (true && false))
	fmt.Printf("\tture || true:  %v\t\n", (true || true))
	fmt.Printf("\tture || false: %v\t\n", (true || false))
	fmt.Printf("\t!ture:         %v\t\n", (!true))
}
