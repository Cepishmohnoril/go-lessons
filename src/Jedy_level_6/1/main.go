package main

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 42
}

bar bar() (int, string) {
	return 451, "farenheit"
}