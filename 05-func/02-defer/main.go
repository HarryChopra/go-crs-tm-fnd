package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo called")
}

func bar() {
	fmt.Println("bar called")
}

/*
bar called
foo called
*/
