package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ninja level 6 Exercises 3:")
	runnow()
	defer icanwait("First to be defered (He who is defered first runs last)")
	runnow()
	defer icanwait("Second to be defered")
	runnow()
	defer icanwait("Third to be defered")
	runnow()
	fmt.Println("Main is done...")
}

func runnow() {
	fmt.Println("RunNow -- Running")
}

func icanwait(s string) {
	fmt.Println("I Can Wait...", s)
}
