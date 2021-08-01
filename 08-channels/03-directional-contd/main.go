package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	wg.Add(1)
	go foo(c) // Does assign bi-directional to directional
	bar(c)    // Does assign bi-directional to directional
	wg.Wait()
}

// send
func foo(c chan<- int) {
	c <- 20
	wg.Done()
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}

/*

20
*/
