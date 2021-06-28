package main

import "fmt"

func main() {
	james := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	james.speak()
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}
