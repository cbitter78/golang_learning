package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Ninja level 9 Exercises 3: Race")
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	var wg sync.WaitGroup
	var tic int64

	// By using an Anonymous function to wrap the go routine I dont
	// need to make the wait group package scope.  All the wait group
	// calls can be made here.  I also found that I can not add to the
	// wait group within the anonymous function block.  It just does
	// not register.  Not sure why..
	wg.Add(1)
	go func() {
		prefix := "go1"
		times := 10
		fmt.Println(prefix, "Starting - Running", times, "times")
		for i := 0; i < times; i++ {
			fmt.Println(prefix, i, "tic:", tic)
			tic++
			runtime.Gosched()
		}
		fmt.Println(prefix, "Done")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		prefix := "go2"
		times := 50
		fmt.Println(prefix, "Starting - Running", times, "times")
		for i := 0; i < times; i++ {
			fmt.Println(prefix, i, "tic:", tic)
			tic++
			runtime.Gosched()
		}
		fmt.Println(prefix, "Done")
		wg.Done()
	}()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}
