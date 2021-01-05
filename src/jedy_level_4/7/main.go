package main

import "fmt"

func main() {
	person1 := []string{"James", "Bond", "Shaken, not stirred"}
	person2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	x := [][]string{person1, person2}

	for _, v := range x {
		for _, vin := range v {
			fmt.Println(vin)
		}
	}
}
