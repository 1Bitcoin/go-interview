package main

import "fmt"

func main() {
	arr := [...]int64{1, 2, 3, 4}

	fmt.Println(arr)

	sl := arr[:]
	sl[0] = 228
	fmt.Println(arr)

	fmt.Println(sl)
}
