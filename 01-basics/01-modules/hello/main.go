package main

import (
	"fmt"

	"example.com/go-crs-tm-fnd/01-basics/01-modules/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}

/*
For production use, we'd publish the example.com/greetings module from its repository
(with a module path that reflected its published location), where Go tools could
find it to download it. For now, because we haven't published the module yet,
we need to adapt the example.com/hello module so it can find the
example.com/greetings code on our local file system:

$ go mod edit -replace example.com/greetings=../greetings

Updates go.mod to require the greetings module version, and downloads source code
into the module cache:

$ go tidy
*/
