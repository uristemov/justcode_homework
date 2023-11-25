package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	cache   = make(map[string]cacheItem)
	mutex   sync.RWMutex
	ticker  *time.Ticker
	cleanup = make(chan struct{})
)

type cacheItem struct {
	Value          interface{}
	ExpirationTime time.Time
}

func run(cleanupInterval time.Duration) {
	ticker = time.NewTicker(cleanupInterval)
	for {
		select {
		case <-ticker.C:
			Expired()
		case <-cleanup:
			ticker.Stop()
			return
		}
	}
}

func Expired() {
	mutex.Lock()
	defer mutex.Unlock()

	now := time.Now()
	for key, item := range cache {
		if now.After(item.ExpirationTime) {
			delete(cache, key)
		}
	}
}

func Set(key string, value interface{}, expiration time.Duration) {
	mutex.Lock()
	defer mutex.Unlock()

	expirationTime := time.Now().Add(expiration)
	cache[key] = cacheItem{Value: value, ExpirationTime: expirationTime}
}

func Get(key string) (interface{}, bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	item, ok := cache[key]
	if !ok {
		return nil, false
	}

	return item.Value, true
}

func Stop() {
	close(cleanup)
}

func main() {
	go run(5 * time.Second)

	Set("key1", "value1", 10*time.Second)
	Set("key2", "value2", 5*time.Second)

	val, exists := Get("key1")
	if exists {
		fmt.Println("Value for key1:", val)
	} else {
		fmt.Println("Key1 not found in cache")
	}

	time.Sleep(8 * time.Second)

	val, exists = Get("key1")
	if exists {
		fmt.Println("Value for key1 after expiration:", val)
	} else {
		fmt.Println("Key1 not found in cache after expiration")
	}

	Stop()
}
