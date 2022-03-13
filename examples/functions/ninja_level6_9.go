package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ninja level 6 Exercises 9:")
	s := []string{"bob", "george", "fread", "susan", "tom"}
	c := func(data string) {
		fmt.Println("Call back with", data)
	}
	processData(c, s)
	processData(otherCallBack, s)
}

func processData(callBack func(string), dataSet []string) {
	fmt.Printf("Starting data processing on %v\n", dataSet)
	for _, v := range dataSet {
		callBack(v)
	}
	fmt.Println("Done Processing data")
}

func otherCallBack(data string) {
	fmt.Println("Other Call Back", data)
}
