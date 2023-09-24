package main

import (
	"fmt"
	"sync"
)

func mergeChannels(channels ...<-chan int) <-chan int {
	mergedCh := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(len(channels))

		for _, channel := range channels {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for num := range ch {
					mergedCh <- num
				}
			}(channel, wg)
		}

		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	d := make(chan int)

	go getChannel(a, []int{1, 2, 3})
	go getChannel(b, []int{4, 5, 6})
	go getChannel(c, []int{7, 8, 9})
	go getChannel(d, []int{10, 11, 12})

	for num := range mergeChannels(a, b, c, d) {
		fmt.Println(num)
	}
}

func getChannel(channel chan<- int, array []int) chan<- int {
	for _, num := range array {
		channel <- num
	}
	close(channel)

	return channel
}
