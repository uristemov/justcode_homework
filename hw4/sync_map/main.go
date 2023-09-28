package main

import (
	"fmt"
	"sync"
)

var data sync.Map
var mutex sync.Mutex

//type SafeMap struct {
//	data  sync.Map
//	mutex sync.Mutex
//}

func Store(key int, value interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	data.Store(key, value)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i <= 10; i++ {
		value := fmt.Sprintf("value %d", i)
		go func(k int, v string) {
			Store(k, v)
			wg.Done()
		}(i, value)
	}

	wg.Wait()

	// Retrieve and print values
	data.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %d, Value: %v\n", key, value)
		return true
	})
}
