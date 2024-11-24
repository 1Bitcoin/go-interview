package main

import "fmt"

func main() {
	slice := []string{"A", "B", "C"}
	addElem(slice[1:2], "X")
	fmt.Println(slice)
}

func addElem(slice []string, elem string) {
	slice = append(slice, elem)
}
