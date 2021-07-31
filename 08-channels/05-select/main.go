package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)
	receive(even, odd, quit)
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- true
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even chan:", v)
		case v := <-o:
			fmt.Println("from the odd chan:", v)
		case <-q:
			fmt.Println("from the quit chan. about to exit")
			return
		}
	}
}
