package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

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
	runtime.Gosched()
	*i++
	wg.Done()
}
