package main

import (
	"fmt"
)

func main() {
	anon()
	buff()
}

func anon() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func buff() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
