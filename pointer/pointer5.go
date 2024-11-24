package main

import "fmt"

type Count int

func (c *Count) inc() {
	*c++
}

func main() {
	var c Count

	c.inc()
	fmt.Println(c)
}
