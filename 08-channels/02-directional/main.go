package main

import "fmt"

func main() {
	c := make(chan int)    // Bidirectional
	cr := make(<-chan int) // Receive channel
	cs := make(chan<- int) // Send Channel

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	// Can't assign a Directional channel to Bidirectional
	// c = cr
	// c = cr

	// Can assign a Bidirectional channel to Directional
	cr = c
	cs = c

	// Similarly, Bidirectional to directional type convert:
	cr1 := (<-chan int)(c)
	cs1 := (chan<- int)(c)

	fmt.Printf("%T\n", cr1)
	fmt.Printf("%T\n", cs1)
}

/*
chan int
<-chan int
chan<- int

<-chan int
chan<- int
*/
