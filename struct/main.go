package main

import "fmt"

type person struct {
	first string
	last  string
}

type emb struct {
	person
	test bool
}

func main() {

	p1 := person{
		first: "Foo",
		last:  "Bar",
	}

	p2 := person{
		first: "Moo",
		last:  "Foo",
	}

	emb1 := emb{
		person: person{
			first: "Hoo",
			last:  "Rar",
		},
		test: true,
	}

	fmt.Println(p1)
	fmt.Println(p2.first)
	fmt.Println(emb1)

	//Same shit
	fmt.Println(emb1.first)
	fmt.Println(emb1.person.first)
}
