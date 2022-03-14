package main

import (
	"encoding/json"
	"fmt"
)

type Dog struct {
	Name         string `json:"Name"`
	Breed        string `json:"Breed"`
	Age          int    `json:"Age"`
	PackLocation string `json:"PackLocation"`
	GoodDog      bool   `json:"GoodDog"`
}

func main() {
	fmt.Println("Ninja level 8 Exercises 2:")
	unmarshel_example()
}

func unmarshel_example() {
	jsonString := `
		[
		{
			"Name": "Hugo",
			"Breed": "Sheppard Pit Mix",
			"Age": 35,
			"PackLocation": "Middle",
			"GoodDog": true
		},
		{
			"Name": "Penny",
			"Breed": "BullDog Pit Mix",
			"Age": 21,
			"PackLocation": "Front",
			"GoodDog": true
		}
		]`
	var dogs []Dog
	err := json.Unmarshal([]byte(jsonString), &dogs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Unmarshel (from json to object)")
	fmt.Println("OBJECT", dogs)
}
