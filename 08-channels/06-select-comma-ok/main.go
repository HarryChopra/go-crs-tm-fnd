package main

import "fmt"

func main() {
	odd := make(chan int)
	even := make(chan int)
	quit := make(chan bool)
	go send(odd, even, quit)
	receive(odd, even, quit)
}

func send(o, e chan<- int, q chan<- bool) {
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(o)
	close(e)
	q <- true
	close(q)
}

// Implementation 1:
// Reading using select on a close channel yields the zero value of the chan type
// refer to the output
// func receive(o, e <-chan int, q <-chan bool) {
// 	for {
// 		select {
// 		case v := <-o:
// 			fmt.Println("received from odd:", v)
// 		case v := <-e:
// 			fmt.Println("received from even:", v)
// 		case v := <-q:
// 			fmt.Println("received from quit:", v)
// 			return
// 		}
// 	}
// }

// Implementation 2: (With comma ok)
// If the chan we're reading from, is closed, 'v' will yield zero value but 'ok' will
// yield false too to indicate that the channel is closed.
// refer to the output
func receive(o, e <-chan int, q <-chan bool) {
	for {
		select {
		case v, ok := <-o:
			if ok {
				fmt.Println("received from odd:", v)
			}
		case v, ok := <-e:
			if ok {
				fmt.Println("received from even:", v)
			}
		case v, ok := <-q:
			if ok {
				fmt.Println("received from quit:", v)
				return
			}
		}
	}
}

/*
output (Implementation 1):

received from odd: 1
received from even: 2
received from odd: 3
received from even: 4
received from odd: 5
received from odd: 0
received from even: 0
received from quit: false

output (Implementation 2):

received from odd: 1
received from even: 2
received from odd: 3
received from even: 4
received from odd: 5
received from quit: true
*/
