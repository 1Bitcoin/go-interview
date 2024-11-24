package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("world")
	}()
	fmt.Println("hello")
	panic("error")
	// hello world error
}
