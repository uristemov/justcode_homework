package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup
var mutex sync.Mutex

func increment() {
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	go increment()
	go increment()

	wg.Wait()

	fmt.Println("Counter:", counter)
}
