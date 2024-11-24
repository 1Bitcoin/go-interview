package main

import "fmt"

func main() {
	var a []int64
	fmt.Println(len(a), cap(a))

	a = append(a, 1)
	fmt.Println(len(a), cap(a))

	a = append(a, 1)
	fmt.Println(len(a), cap(a))

	a = append(a, 1)
	fmt.Println(len(a), cap(a))

	a = append(a, 1)
	fmt.Println(len(a), cap(a))

}
