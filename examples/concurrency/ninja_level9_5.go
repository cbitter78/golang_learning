package main

// See https://gobyexample.com/atomic-counters
// See https://github.com/GoesToEleven/golang-web-dev/blob/master/000_temp/52-race-condition/05_atomic/main.go

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Ninja level 9 Exercises 4: Race with Atomic")
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter uint64

	const gs = 113
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := atomic.AddUint64(&counter, 1)
			runtime.Gosched()
			fmt.Println("\tCounter:", v)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
