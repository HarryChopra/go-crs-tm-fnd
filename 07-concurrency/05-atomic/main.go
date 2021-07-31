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
			atomic.AddInt64(&counter, 1)
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
