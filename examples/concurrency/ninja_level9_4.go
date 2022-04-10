package main

// See https://gobyexample.com/mutexes
// See https://github.com/GoesToEleven/golang-web-dev/tree/master/000_temp/52-race-condition/04_mutex

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Ninja level 9 Exercises 4: Race with Mutex")
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			v := counter
			v++
			counter = v
			fmt.Println("\tCounter:", counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
