package main

import "fmt"

func main() {
	a := 2

	defer func() {
		test3()()
	}()

	a++

	fmt.Println("end")
	fmt.Println(a)
}

func test1(a int) {
	fmt.Println(a)
}

func test2(a int) int {
	return a * a
}

func test3() func() {
	a := 3
	return func() {
		fmt.Println(a * a)
	}
}
