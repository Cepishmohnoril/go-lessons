package main

import "fmt"

func main() {
	ch := make(chan int)

	go put(ch)
	get(ch)

	fmt.Println("That's all folks")
}

func put(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func get(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
