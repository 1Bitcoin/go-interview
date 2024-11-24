package main

import (
	"fmt"
	"log"
)

func handle1(s1 []int) {
	s1[0] = 111
}
func handle2(s1 []int) []int {
	return append(s1, 12345)
}

func main() {
	s1 := []int{0, 1, 2} // len=3 cap=3

	fmt.Println(len(s1), cap(s1))

	log.Printf("before: %v", s1) // 0 1 2

	handle1(s1)
	log.Printf("after handle1: %v", s1) // что будет напечатано? // 111 1 2

	s1 = handle2(s1)
	log.Printf("after handle2: %v", s1) // что будет напечатано? // 111 1 2
}
