package main

import (
	"bytes"
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
	fmt.Println("Ninja level 8 Exercises 1:")
	marshel_example()
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
