package main

import "fmt"

func main() {
	foo := `this is raw string literal \n you see, it doesen't "work"`
	fmt.Println(foo)
}
