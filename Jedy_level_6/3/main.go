package main

import "fmt"

func main() {
	defer first()
	second()
}

func first() {
	fmt.Println("First function")
}

func second() {
	fmt.Println("Second function")
}
