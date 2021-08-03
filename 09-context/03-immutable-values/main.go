package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Printf("doSomething: Value at myKey = %s\n", ctx.Value("myKey"))
	// Context's data is immutable. WithValue returns a new context.
	anotherCtx := context.WithValue(ctx, "myKey", "anotherVal")
	doAnotherThing(anotherCtx)
	fmt.Printf("doSomething: Value at myKey = %s\n", ctx.Value("myKey"))
}

func doAnotherThing(ctx context.Context) {
	fmt.Printf("doAnotherThing: Value at myKey = %s\n", ctx.Value("myKey"))
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myVal")
	doSomething(ctx)
}
