package main

import "fmt"

func main() {
	a := make(chan int, 1)
	for {
		select {
		case val := <-a:
			fmt.Println(val)
		}
	}
}
