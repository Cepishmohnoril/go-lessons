package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x) // & gives address

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", &x)

	z := &x
	fmt.Println(z)
	fmt.Println(*z) // * gives value from addres
}
