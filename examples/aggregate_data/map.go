package main

import (
	"fmt"
)

func main() {

	fmt.Println("Lets play with maps!")

	fmt.Println("Lets create a map and add some data to it")
	m := map[string]int{
		"bob": 9,
	}
	m["foo"] = 1
	m["bar"] = 10
	fmt.Println(m)

	v, ok := m["bob"]
	fmt.Printf("v:%v ok:%v\n", v, ok)

	// Check to see if the key is in the map and use it on one line.  :)
	if v, ok := m["foo"]; ok {
		fmt.Printf("In Map - v:%v ok:%v\n", v, ok)
	} else {
		fmt.Printf("Not in Map - v:%v ok:%v\n", v, ok)
	}

	println("Lets iterate over the map with range")
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Remove the foo entry
	delete(m, "foo")
	delete(m, "not there!") // This is a noop and no error is thrown

	println("Lets iterate over just the values in the map")
	for _, v := range m {
		fmt.Println(v)
	}

}
