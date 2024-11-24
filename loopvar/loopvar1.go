package main

import (
	"fmt"
)

func main() {
	numbers := []*int{}
	for i := 0; i < 5; i++ {
		numbers = append(numbers, &i)
	}
	for _, number := range numbers {
		fmt.Printf("%d ", *number) // что отобразится после вызова?
	}
	// 5 5 5 5 5 go <1.22
	// 0 1 2 3 4  go >1.22
}
