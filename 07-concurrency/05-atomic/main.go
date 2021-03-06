package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	nroutines := 100
	var wg sync.WaitGroup

	wg.Add(nroutines)
	for i := 0; i < nroutines; i++ {
		go func() {
			// Read/Write avoiding race (locking and unlocking)
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

/*
go run -race main.go
Counter: 100
*/
