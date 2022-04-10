package main

import (
	"fmt"
	"math/rand"
	"time"
)

type teenager struct {
	name          string
	age           int
	maturityLevel int
	favoriteColor string
	collageChoice string
}

type kid struct {
	name          string
	age           int
	favoriteColor string
}

// The method set below for *teenager attach to the pointer to teenager
// not the type teenager.  Therefor the type does not implment
// the interfaces below.  But the pointer to it `*teenager` does.
// when working with a variable of type teenager it will use methods
// that recieve *teenager and teenager. It dont care.  Interfaces care!
func (t *teenager) doWhatISay(ask string) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(100000000)
	if r%2 == 0 { // Implement teenager logic
		fmt.Println("Teenager:\t", t.name, "refused your ask:", ask, "\t\t:^(")
	} else {
		fmt.Println("Teenager:\t", t.name, "Will do what you asked:", ask, "\t:^)")
	}
}
func (t *teenager) String() string {
	return fmt.Sprintf("Teenager: \t%v (%v)\tMaturity Level:  %v\t Collage: %v", t.name, t.age, t.maturityLevel, t.collageChoice)
}

func (k kid) doWhatISay(ask string) {
	fmt.Println("Kid:\t\t", k.name, "Will do waht you asked", ask, "\t:)")
}
func (k kid) String() string {
	return fmt.Sprintf("Kid:\t\t%v (%v)\tFavorite Color:  %v", k.name, k.age, k.favoriteColor)
}

type ObeyMe interface {
	doWhatISay(ask string)
}
type SoundOff interface {
	String() string
}

func cleanYourRoom(o ObeyMe) {
	o.doWhatISay("Clean your room")
}
func doTheDishes(o ObeyMe) {
	o.doWhatISay("Do the Dishes")
}
func whosThere(s SoundOff) {
	fmt.Println("Sound Off:\t", s.String())
}

func main() {
	fmt.Println("Ninja level 9 Exercises 2:")

	tony := teenager{
		name:          "Tony",
		age:           19,
		favoriteColor: "blue",
		collageChoice: "TBD",
		maturityLevel: 6,
	}
	luke := kid{
		name:          "Luke",
		age:           16,
		favoriteColor: "green",
	}

	fmt.Println("----", "Calling method sets directly", "----")
	fmt.Println("Who are you?", tony.String())
	tony.doWhatISay("Stand on your head")

	fmt.Println("----", "Calling Interfaces", "----")
	whosThere(luke)
	whosThere(&tony)

	doTheDishes(luke)
	doTheDishes(&tony)

	cleanYourRoom(luke)
	doTheDishes(&tony)

}
