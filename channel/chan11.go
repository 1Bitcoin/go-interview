package main

import "fmt"

func main() {

	in := make(chan int, 5)

	go func() {
		in <- 1
		in <- 2
		in <- 3
		close(in)
	}()

	for {
		select {
		case value, ok := <-in:
			if !ok {
				fmt.Println("goodbye")
				return
			}

			fmt.Println(value)
		}
	}
}
