package main

import "fmt"

func main() {
	//	init; condition; post;
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	x := 1
	for x <= 10 {
		fmt.Println(x)
		x++
	}

	/*
		for {
			Infinite loop. Cool!
		}
	*/

	if 2 == 2 {
		fmt.Println("Herro")
	}

	if x := 10; x == 42 {
		fmt.Println("No Herro")
	}

	//x left inside if scope

	if false {
		//not do
	} else if true {
		fmt.Println("Test")
	} else {
		//no do
	}

}
