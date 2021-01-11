package main

import "fmt"

type person struct {
	firstName       string
	lastname        string
	favoriteFlavors []string
}

func main() {
	p1 := person{
		firstName:       "Hello",
		lastname:        "World",
		favoriteFlavors: []string{"Cherry", "Chocolate"},
	}

	p2 := person{
		firstName:       "Foo",
		lastname:        "Bar",
		favoriteFlavors: []string{"Cherry", "Vanilla"},
	}

	fmt.Println(p1)
	for _, flavor := range p2.favoriteFlavors {
		fmt.Println(flavor)
	}
}
