package main

import "fmt"

func main() {
	c := make(chan int)
	go send(c)
	// Iterate over channel and receive its messages
	for v := range c {
		fmt.Println(v)
	}
}

func send(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	// Close the channel after the last message is sent,
	// or the caller routine will go into a deadlock
	close(c)
}
