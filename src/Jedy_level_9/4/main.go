package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	sum := 0
	goal := 100

	wg.Add(goal)
	for i := 0; i < goal; i++ {
		go add(&sum)
	}

	wg.Wait()

	fmt.Println(sum)
}

func add(i *int) {
	mu.Lock()
	*i++
	mu.Unlock()
	wg.Done()
}
