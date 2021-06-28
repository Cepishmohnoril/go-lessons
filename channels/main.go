package main

import "fmt"

func main() {
	basicChannel()
	directionalChannel()
	rangeChannel()
}

func basicChannel() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	fmt.Println(<-ch)
}

func directionalChannel() {

	ch := make(chan int)

	go send(ch)

	recive(ch)
}

func send(c chan<- int) {
	c <- 42
}

func recive(c <-chan int) {
	fmt.Println(<-c)
}

func rangeChannel() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}
}
