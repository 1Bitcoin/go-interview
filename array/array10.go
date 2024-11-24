package main

import "log"

func handlel(s1 [3]int) {
	s1[0] = 111
}

func main() {
	s1 := [...]int{0, 1, 2}

	log.Printf("before: %v", s1)

	handlel(s1)
	log.Printf("after handle1: %v", s1) // что будет напечатано?
}
