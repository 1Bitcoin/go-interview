package main

import (
	"fmt"
)

func main() {
	a := 10
	defer func(a int) {
		fmt.Println("call 0 ", a+10)
	}(a)

	defer func() {
		fmt.Println("call 1 ", a+10)
	}()

	defer fmt.Println("call 2 ", a+10)

	a++

	fmt.Println("call 3 ", a)
}

// call 2 11
// call 1 21
// call 0 20
