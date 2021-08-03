package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	deadline := time.Now().Add(time.Duration(200 * time.Millisecond))
	ctx, cancelCtx := context.WithDeadline(ctx, deadline)
	printCh := make(chan int)
	go doAnotherThing(ctx, printCh)
loop:
	for i := 1; i <= 3; i++ {
		select {
		case <-ctx.Done():
			break loop
		case printCh <- i:
			time.Sleep(100 * time.Millisecond)
		}
	}

	cancelCtx()
	fmt.Println("doSomething: finished")
}

func doAnotherThing(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnotherThing: err = %s\n", err)
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
