package main

import (
	"fmt"
)

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func checkErr(err error) {
	fmt.Println(err == nil)
}

func main() {
	var e1 error
	checkErr(e1) // что отобразится после вызова? // true

	var e *errorString
	checkErr(e) // что отобразится после вызова? // false

	e = &errorString{}
	checkErr(e) // что отобразится после вызова? // false

	e = nil
	checkErr(e) // что отобразится после вызова? // false
}
