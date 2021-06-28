package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("This is callback")
	}

	call(f)
}

func call(f func()) {
	f()
}
