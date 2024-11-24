package main

import (
	"fmt"
)

func pipeline() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x <= 10; x++ {
			naturals <- x
		}
		defer close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		defer close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
