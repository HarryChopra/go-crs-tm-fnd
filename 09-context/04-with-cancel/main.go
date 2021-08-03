package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)
	printCh := make(chan int)
	go doAnotherThing(ctx, printCh)
	for i := 1; i <= 3; i++ {
		printCh <- i
	}
	cancelCtx()
	time.Sleep(100 * time.Millisecond) // Optional: To print rest of output (cancel)
	fmt.Println("doSomething: finished")

}

func doAnotherThing(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnotherThing: err = %q\n", err)
			}
			fmt.Println("doAnotherThing: finished")
			return

		case num := <-ch:
			fmt.Printf("doAnotherThing: %d\n", num)
		}
	}
}

func main() {
	ctx := context.Background()
	doSomething(ctx)
}
