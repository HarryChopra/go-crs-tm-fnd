package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var wg sync.WaitGroup
	var total int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				// runtime.Gosched()
				total += getRand()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("total:", total)
}

//go:noinline
func getRand() int {
	return rand.Intn(10)
}

/*
Go implements asynchronous preemption from Go 1.14.
As a result, loops without function calls no longer potentially deadlock the scheduler
or significantly delay garbage collection.
Any goroutine running for more than 10ms is marked as preemptible.
*/
