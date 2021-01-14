package main

import "fmt"

func main() {
	x := getFunc()
	x()
}

func getFunc() func() {
	return func() {
		fmt.Println("Hello function from another function")
	}
}
