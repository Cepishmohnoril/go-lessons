package main

import "fmt"

func main() {

	foo := "bar"

	fmt.Printf("%v the address of value `%v`", &foo, foo)
}
