package main

import (
	"fmt"
)

func run1() {
	ch := make(chan string)
	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
		}
	}()
	ch <- "cmd.1"
	ch <- "cmd.2"
}

func main() {
	run1()
}
