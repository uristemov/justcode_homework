package main

import "fmt"

func channel_num(f func(int) int, number chan int) {

	number <- f(10)

}

func main() {
	num := make(chan int) //initializing the channel

	go channel_num(square, num) //calling as a goroutine

	i := <-num
	j := <-num // waiting num from chan but it will never come DEADLOCK
	fmt.Println("Value of Channel i,j =", i, j)
}

func square(num int) int {
	return num * num
}
