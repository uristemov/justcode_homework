package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		wg.Done()
		fmt.Println("goroutine on process")
	}(wg)
	wg.Wait()
	<-ch                //No sender
	fmt.Println("end.") //Will never reach
}
