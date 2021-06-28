package main

import (
	"fmt"
)

type customErr struct {
}

func (ce customErr) Error() string {
	return "Some custom error"
}

func main() {
	c1 := customErr{}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
