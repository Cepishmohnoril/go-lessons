package main

import "fmt"

func main() {
	fmt.Println(Sum(2, 3))
}

//Sum add all supplied values and returns int
func Sum(xi ...int) int {
	sum := 0
	for _, val := range xi {
		sum += val
	}

	return sum
}
