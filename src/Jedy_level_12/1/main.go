package main

import (
	"dog"
	"fmt"
)

func main() {
	humanYears := 5
	dogYears := dog.Years(humanYears)

	fmt.Println(humanYears, "human years are", dogYears, "dog years")
}
