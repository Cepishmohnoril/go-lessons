package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {

	orig := person{
		first: "Foo",
		last:  "Bar",
	}

	changeMe(&orig)

	fmt.Println(orig)
}

func changeMe(p *person) {
	(*p).last = "Boop"
}
