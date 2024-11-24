package main

import "fmt"

func main() {
	fmt.Println("start")

	for _, v := range []int{1, 2, 3, 4, 5} {
		defer fmt.Println(v)
	}

	fmt.Println("end")
}
