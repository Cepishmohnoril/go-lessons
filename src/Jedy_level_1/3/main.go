package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%d %v %t", x, y, z)
	fmt.Println(s)
}
