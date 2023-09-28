package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup

func increment() {
	for i := 0; i < 1000; i++ {
		counter++
	}
	wg.Done()
}

func main() { //go run -race main.go
	wg.Add(2)

	go increment()
	go increment()

	wg.Wait()

	fmt.Println(counter)
}
