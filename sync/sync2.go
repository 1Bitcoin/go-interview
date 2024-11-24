package main

import (
	"fmt"
	"sync"
	"time"
)

var m map[string]int

func main() {
	m = make(map[string]int)
	go f1()
	go f2()

	time.Sleep(time.Second * 1)

	fmt.Printf("print map = %+v", m)
}

func f1() {
	for i := 0; i < 100000; i++ {
		m["f1"]++
	}
}

func f2() {
	for i := 0; i < 100000; i++ {
		m["f2"]++
	}
}
