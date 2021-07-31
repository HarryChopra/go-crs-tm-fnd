package main

import (
	"fmt"
	"sync"
)

// // Case 1: c blocks the main routine
// // As chan c is unbuffered, main will wait, endlessly, to write to c -
// // as there is no other routine to read from it.
// // Deadlock
// func main() {
// 	var wg sync.WaitGroup
// 	c := make(chan int)

// 	c <- 1
// }

//
// // Case 2: c blocks the anonymous Go routine
// // Deadlock, Go Routine waiting to read from c, as there is no other routine to write to it.
// // main is ultimately forced to wait by the waitgroup wg.
// func main() {
// 	var wg sync.WaitGroup
// 	c := make(chan int)

// 	wg.Add(1)
// 	go func() {
// 		fmt.Println(<-c)
// 		wg.Done()
// 	}()
// 	wg.Wait()
// }

//
// Case 2: c blocks the anonymous Go routine
// 4 send - 2 buffer - 1 read = 1 waiting to send
// Go routine endlessly waits to write to c, as there is no other routine to read from it.
// main is ultimately forced to wait by the waitgroup wg.
// Deadlock
func main() {
	var wg sync.WaitGroup
	c := make(chan int, 2)

	wg.Add(2)
	go func() {
		c <- 1
		wg.Done()
	}()
	go func() {
		c <- 1
		c <- 1
		c <- 1
		wg.Done()
	}()
	fmt.Println(<-c)
	wg.Wait()
}
