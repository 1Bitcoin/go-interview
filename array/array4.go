package main

import "fmt"

func main() {
	a := []int{1, 2, 3} // len=3 cap=3
	b := a

	b = append(b, 4) // 1 2 3 4 [b] len=4 cap=6
	c := b           // 1 2 3 4 [c] len=4 cap=6

	b[0] = 0          // 0 2 3 4 [b] len=4 cap=6
	e := append(c, 5) // 0 2 3 4 5 [b] len=5 cap=6
	b[2] = 7          // 0 2 7 4 5 [b] len=5 cap=6

	fmt.Println(a, b, c, e) // что отобразится после вызова?
	// [1 2 3] [0 2 7 4] [0 2 7 4] [0 2 7 4 5]
}
