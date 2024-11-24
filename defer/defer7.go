package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5}

	for i := range arr {
		defer fmt.Printf("%d ", arr[i])
	}
}
