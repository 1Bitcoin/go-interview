package main

import (
	"fmt"
)

func testSlices2() {
	a := []byte{'a', 'b', 'c'} // len=3 cap=3
	b := append(a[1:2], 'd')   // b d len=2 cap=3
	b[0] = 'z'                 // z d len=2 cap=3

	fmt.Printf("%s\n", a) // что отобразится после вызова?
	// a z d
}

func main() {
	testSlices2()
}
