package main

import (
	"fmt"
)

func main() {
	in := make(chan int)
	out := make(chan int)

	go work1(in)
	go work2(in, out)

	for value := range out {
		fmt.Println(value)
	}
}

func work1(in chan int) {

	for value := range []int64{1, 2, 3, 4, 5, 6} {
		in <- value
	}
	close(in)
}

func work2(in, out chan int) {

	for value := range in {
		out <- value * value
	}
	close(out)
}
