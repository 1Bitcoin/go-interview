package main

import "fmt"

func main() {
	var m map[string]int

	fmt.Println(m["foo"])

	m["foo"] = 42

	fmt.Println(m["foo"])
}
