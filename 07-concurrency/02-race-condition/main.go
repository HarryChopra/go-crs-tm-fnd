package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var counter int
	routines := 100
	var wg sync.WaitGroup

	wg.Add(routines)
	for i := 0; i < routines; i++ {
		go func() {
			v := counter
			v++
			time.Sleep(time.Millisecond)
			counter = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}

/*
go run -race main.go

Counter: 7
Found 1 data race(s)
*/
