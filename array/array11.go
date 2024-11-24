package main

import "fmt"

func main() {
	a := make([]int, 0, 3)
	a = append(a, 1)
	a = append(a, 2)
	add(a)
	fmt.Printf("%v\n", a)
}

func add(a []int) {
	a = append(a, 3)
	a = append(a, 4)
}
