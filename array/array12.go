package main

import "fmt"

func main() {
	s := "test"
	println(rune(s[0]))
	fmt.Println(string(s[0])) // Печатаем первый символ (код первого байта)
	//s[0] = 'A'
	//println(s)
}
