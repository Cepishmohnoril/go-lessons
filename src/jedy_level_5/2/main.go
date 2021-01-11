package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	favoriteFlavors []string
}

func main() {
	p1 := person{
		firstName:       "James",
		lastName:        "Bond",
		favoriteFlavors: []string{"Cherry", "Chocolate"},
	}

	p2 := person{
		firstName:       "Dog",
		lastName:        "Snoop",
		favoriteFlavors: []string{"Cherry", "Vanilla"},
	}

	persons := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range persons {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for _, flavor := range v.favoriteFlavors {
			fmt.Println(flavor)
		}

		fmt.Println("------")
	}
}
