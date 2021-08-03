package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Printf("doSomething: Value of myKey = %s\n", ctx.Value("myKey"))
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myVal")
	doSomething(ctx)
}
