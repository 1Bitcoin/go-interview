package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	Sex   bool
	Arr   [6]byte
	Gan   float64
	empty struct{}
}

func main() {
	//a := A{
	//	Sex: false,
	//	Gan: 22,
	//	Arr: [6]byte{},
	//	empty: struct{}{},
	//}
	//fmt.Println(unsafe.Sizeof(a))

	v := int64(1)
	fmt.Println(unsafe.Sizeof(&v))

}
