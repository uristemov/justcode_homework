package main

import (
	"fmt"
	"sync"
)

var data = make(map[int]string)
var rwMutex sync.RWMutex
var wg sync.WaitGroup

func readData(key int) string {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	wg.Done()
	return data[key]
}

func writeData(key int, value string) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	data[key] = value
	wg.Done()
}

func main() {

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		value := fmt.Sprintf("value %d", i)
		go writeData(i, value)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go readData(i)
	}

	wg.Wait()

	fmt.Println("Map:", data)
}
