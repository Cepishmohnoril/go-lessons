package main

import "fmt"

func main() {
	foo()
	bar("Doot")
	fmt.Println(woo("Doomslayer"))

	x, y := tang("He", "hellwalker")

	fmt.Println(x, y)

	variadic(1, 2, 3, 4, 5, 6)

	sl1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sum(sl1...))
	fmt.Println(sum())

	def()
	method()

	//Anonymous function

	func() {
		fmt.Println("Anonymous function run")
	}()

	func(x int) {
		fmt.Println("The meaning of life is:", 42)
	}(42)

	//func expression

	f := func(x int) {
		fmt.Println("The number is:", x)
	}

	f(42)

	//return function

	f2 := returnFunc()

	fmt.Println(f2(451))

	//callback func
	passFunc(func() {
		fmt.Println("This is a callback call.")
	})
}

func foo() {
	fmt.Printf("Hello from foo.\n")
}

func bar(name string) {
	fmt.Println("Hello,", name)
}

func woo(name string) string {
	return fmt.Sprintln("Hello from woo,", name)
}

func tang(person string, desc string) (string, bool) {
	phrase := fmt.Sprintln(person, "is a", desc)
	return phrase, true
}

func variadic(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}

//Defer example

func def() {
	defer foo()
	bar("this one goes first")
}

//Method example

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("My name is", s.fname, s.lname)
}

type human interface {
	speak()
}

func method() {
	sa1 := secretAgent{
		person: person{
			fname: "James",
			lname: "Bond",
		},
		ltk: true,
	}

	sa1.speak()

	inter(sa1)
}

func inter(hum human) {
	hum.speak()
	fmt.Println("I am a human.")
}

func returnFunc() func(int) string {
	return func(x int) string {
		return fmt.Sprintln("Temperature in Farenheits", x)
	}
}

func passFunc(f func()) {
	f()
}
