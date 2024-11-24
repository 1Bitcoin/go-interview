package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)

	select {
	case ch1 <- 1:
		fmt.Println("test")
	}
}
