package main

import (
	"sync"
	"time"
)

func worker1() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
	}()

	return ch
}

func main() {
	timeStart := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		_ = <-worker1()

	}()

	go func() {
		defer wg.Done()
		_ = <-worker1()

	}()

	wg.Wait()

	println(int(time.Since(timeStart).Seconds())) // что выведет - 3 или 6?
}
