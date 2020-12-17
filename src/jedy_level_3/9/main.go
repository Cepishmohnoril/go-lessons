package main

import "fmt"

func main() {
	favSpot := "Two"

	switch favSpot {
	case "One":
		fmt.Println("First spot")
	case "Two":
		fmt.Println("Second spot")
	}
}
