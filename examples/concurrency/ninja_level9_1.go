package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Ninja level 9 Exercises 1: Using WaitGroups")
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	var wg sync.WaitGroup

	// By using an Anonymous function to wrap the go routine I dont
	// need to make the wait group package scope.  All the wait group
	// calls can be made here.  I also found that I can not add to the
	// wait group within the anonymous function block.  It just does
	// not register.  Not sure why..
	wg.Add(1)
	go func() {
		// wg.Add(1)  <-- This DOES NOT Work  It has to go outside
		printTimes("Go Routine One", 10)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		printTimes("Go Routine Two", 47)
		wg.Done()
	}()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func printTimes(prefix string, times int) {
	fmt.Println(prefix, "Starting - Running", times, "times")
	for i := 0; i < times; i++ {
		fmt.Println(prefix, i)
	}
	fmt.Println(prefix, "Done")
}
