package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Ninja level 9 Exercises 6: Print Stuff")
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Arch:\t\t", runtime.GOARCH)
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("Go Routines:\t", runtime.NumGoroutine())
	fmt.Println("Go Root:\t", runtime.GOROOT())

}
