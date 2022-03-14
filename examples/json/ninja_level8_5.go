package main

import (
	"fmt"
	"sort"
	"strings"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []User

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLastName []User

func (a ByLastName) Len() int           { return len(a) }
func (a ByLastName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLastName) Less(i, j int) bool { return a[i].Last < a[j].Last }

/*
String function to make printing easy.
*/
func (u User) String() string {
	b := new(strings.Builder)
	fmt.Fprintf(b, "%v %v Age(%v) Says:\n", u.First, u.Last, u.Age)
	for _, saying := range u.Sayings {
		fmt.Fprintf(b, "\t%v\n", saying)
	}
	fmt.Fprintln(b)
	return b.String()
}

func main() {
	fmt.Println("Ninja level 8 Exercises 5:")
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}
	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}
	for _, u := range users {
		sort.Strings(u.Sayings)
	}

	sort.Sort(ByAge(users))
	fmt.Println(users)

	sort.Sort(ByLastName(users))
	fmt.Println(users)

	// your code goes here

}
