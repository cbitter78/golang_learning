package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Just use https://mholt.github.io/json-to-go/ and feed it json to get
// a moce strict
type Dog struct {
	Name         string `json:"Name"`
	Breed        string `json:"Breed"`
	Age          int    `json:"Age"`
	PackLocation string `json:"PackLocation"`
	GoodDog      bool   `json:"GoodDog"`
}

func main() {
	marshel_example()
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

func marshel_example() {
	d1 := Dog{
		Name:         "Hugo",
		Breed:        "Sheppard Pit Mix",
		Age:          35,
		PackLocation: "Middle",
		GoodDog:      true,
	}
	d2 := Dog{
		Name:         "Penny",
		Breed:        "BullDog Pit Mix",
		Age:          21,
		PackLocation: "Front",
		GoodDog:      true,
	}
	dogs := []Dog{d1, d2}
	fmt.Println("OBJECT:", dogs)
	bs, err := json.Marshal(dogs)
	if err != nil {
		fmt.Println("Error while trying to Marshel to json", err)
		bs = []byte{}
	}

	bs, _ = prettyprint(bs)
	fmt.Println("JSON:\n", string(bs))
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
