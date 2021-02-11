package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go put(ch)
		}

		wg.Wait()
		close(ch)
	}()

	get(ch)

	fmt.Println("That's all folks")
}

func put(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	wg.Done()
}

func get(ch <-chan int) {
	i := 0

	for v := range ch {
		i++
		fmt.Println(i, v)
	}
}
