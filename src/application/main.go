package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	fortunateJSON()
}

func fortunateJSON() {
	p1 := person{
		First: "Foo",
		Last:  "Bar",
		Age:   32,
	}

	p2 := person{
		First: "Hello",
		Last:  "World",
		Age:   42,
	}

	crowd := []person{p1, p2}

	bs, err := json.Marshal(crowd)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	var crowd2 []person

	err2 := json.Unmarshal(bs, &crowd2)

	if err2 != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", crowd2)
}
