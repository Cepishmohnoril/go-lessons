package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var sum int64
var wg sync.WaitGroup

func main() {
	goal := 100

	wg.Add(goal)
	for i := 0; i < goal; i++ {
		go add()
	}

	wg.Wait()

	fmt.Println(sum)
}

func add() {
	atomic.AddInt64(&sum, 1)
	wg.Done()
}
