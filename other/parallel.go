package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	in := make(chan int, 100)
	out := make(chan int, 100)

	// Generating input data
	go func() {
		for i := 0; i < 100; i++ {
			in <- i
		}
		close(in)
	}()

	processInParallel(in, out, 5)

	for result := range out {
		fmt.Println("Processed:", result)
	}
}

func processInParallel(in <-chan int, out chan<- int, workers int) {
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for data := range in {
				out <- processData(data)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

}

func processData(data int) int {
	time.Sleep(100 * time.Millisecond) // Simulate time-consuming task
	return data * 2
}
