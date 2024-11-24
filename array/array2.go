package main

import "fmt"

func main() {
	arr1 := [...]int64{1, 2, 3, 4, 5} // len=5 cap=5
	fmt.Printf("arr1. len=%d, cap=%d\n", len(arr1), cap(arr1))

	arr2 := arr1[:4] // len=4 cap=5
	fmt.Printf("arr2. len=%d, cap=%d\n", len(arr2), cap(arr2))

	arr3 := append(arr2[:2], 9, 8, 7)
	// len=2 cap=5
	// 1 2 9 8 7
	fmt.Printf("arr3. len=%d, cap=%d\n", len(arr3), cap(arr3))

	fmt.Println(arr3)
	fmt.Println(arr2)
	fmt.Println(arr1)

	//[1 2 9 8 7]
	//[1 2 9 8]
	//[1 2 9 8 7]
}
