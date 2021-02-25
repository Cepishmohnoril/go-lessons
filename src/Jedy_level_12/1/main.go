package main

import (
	"fmt"

	"github.com/Cepishmohnoril/go-lessons/src/Jedy_level_12/1/dog"
)

func main() {
	humanYears := 5
	dogYears := dog.Years(humanYears)

	fmt.Println(humanYears, "human years are", dogYears, "dog years")
}
