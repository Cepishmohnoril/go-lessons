package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(arr[:5])
	fmt.Println(arr[5:])
	fmt.Println(arr[2:7])
	fmt.Println(arr[1:6])

}
