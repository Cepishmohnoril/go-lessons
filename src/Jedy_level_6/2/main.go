package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4}

	fmt.Println(foo(x...))

	fmt.Println(bar(x))
}

func foo(x ...int) int {
	var sum int
	for _, digit := range x {
		sum += digit
	}

	return sum
}

func bar(x []int) int {
	return foo(x...)
}
