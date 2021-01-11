package main

import "fmt"

func main() {
	x := struct {
		foo string
		bar bool
	}{
		foo: "Hello World",
		bar: true,
	}

	fmt.Println(x.foo)

}
