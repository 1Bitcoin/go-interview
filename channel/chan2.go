package main

import (
	"fmt"
)

func test(ch chan int64) {
	for i := 0; i <= 3; i++ {
		num := <-ch
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main start")
	ch := make(chan int64, 3)

	go test(ch)

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println("main stop")
}
