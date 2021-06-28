package main

import "fmt"

func main() {
	year := 2020

	switch {
	case year == 1992:
		fmt.Printf("%v was a good year", year)
	case year == 2020:
		fmt.Printf("%v was a bad year", year)
	}
}
