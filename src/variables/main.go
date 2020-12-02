package main

import "fmt"

var y int

type test int

var a test

func main() {
	x := "Hello"
	fmt.Println(x)

	x = "Foo"
	fmt.Println(x)

	fmt.Println(y)

	z := 0.0
	fmt.Printf("%T\n", z)

	a = 5
	fmt.Printf("%T\n", a)
}
