package main

import (
	"os"
)

func main() {
	_, e := os.Open("fooBar.txt")

	if e != nil {
		//fmt.Println(e)
		//log.Println(e)
		//log.Fatal(e)
		panic(e)
	}
}
