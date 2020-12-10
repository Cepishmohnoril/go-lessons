package main

import "fmt"

func main() {
	x := 2

	fmt.Printf("%d\n", x) //dec
	fmt.Printf("%b\n", x) //bin
	fmt.Printf("%x\n", x) //hex

	x = x << 1

	fmt.Printf("%d\n", x) //dec
	fmt.Printf("%b\n", x) //bin
	fmt.Printf("%x\n", x) //hex
}
