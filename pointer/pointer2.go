package main

import (
	"fmt"
)

func main() {
	v := 5
	p := &v
	fmt.Println(*p) // что отобразится после вызова? // 5

	changePointer(p)
	fmt.Println(*p) // что отобразится после вызова? // 5
}

func changePointer(p *int) {
	v := 3
	p = &v
}
