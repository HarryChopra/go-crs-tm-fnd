package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// fmt.Println("OS:\t\t", runtime.GOOS)
	// fmt.Println("ARCH:\t\t", runtime.GOARCH)
	// fmt.Println("CPUs:\t\t", runtime.NumCPU())
	// fmt.Println("GoRoutines:\t", runtime.NumGoroutine())

	wg.Add(1)   // Waitgroup counter
	go func() { // Go routine
		fn("foo")
		wg.Done() // Decrement the counter
	}()
	fn("bar")

	wg.Wait() // Blocks main routine from returning until waitgroup reaches 0
}

func fn(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%s:%d\n", name, i)
		time.Sleep(1 * time.Millisecond)
	}
}
