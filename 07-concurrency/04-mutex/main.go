package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var counter int
	nroutines := 100
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(nroutines)
	for i := 0; i < nroutines; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			time.Sleep(time.Millisecond)
			counter = v
			mu.Unlock()
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
