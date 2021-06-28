package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type human interface {
	speak()
}

func main() {
	hum := person{
		First: "John",
		Last:  "Doe",
		Age:   42,
	}

	saySomething(&hum)
}

func (p *person) speak() {
	fmt.Println("Hello, my name is", p.First)
}

func saySomething(h human) {
	h.speak()
}
