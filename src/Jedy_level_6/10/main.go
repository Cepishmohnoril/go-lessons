package main

import "fmt"

var foo string = "This is foo"

func main() {
	first()
	second()
	third()

	foo := "foo"
	var bar string

	{
		bar := "bar"
		fmt.Println(foo, bar)
	}

	fmt.Println(foo, bar)
}

func first() {
	bar := "This is bar"
	fmt.Println(foo, bar)
}

func second() {
	bar := "This is another bar"
	fmt.Println(foo, bar)
}

func third() {
	foo := "This is another foo"
	fmt.Println(foo)
}
