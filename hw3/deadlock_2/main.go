package main

import "fmt"

func main() {
	nums := make(chan int)
	squares := make(chan int)

	go func() {
		for num := range nums {
			squares <- num * num
		}
		//close(nums) DEADLOCK
	}()

	go func() {
		//кладем значения в канал nums
		for _, num := range []int{1, 2, 3} {
			nums <- num
		}
		//close(nums) DEADLOCK
	}()

	for num := range squares {
		fmt.Println(num)
	}
}
