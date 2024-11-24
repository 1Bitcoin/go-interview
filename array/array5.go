package main

import (
	"fmt"
)

func testSlices1() {
	a := []string{"a", "b", "c"}
	b := a[1:2] // b len=1 cap=3
	b[0] = "q"  // q len=1 cap=3

	fmt.Printf("%s\n", a) // что отобразится после вызова?
	// a q c
}

func main() {
	testSlices1()
}
