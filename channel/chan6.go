package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan string, 3)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(с chan<- string, i int, group *sync.WaitGroup) {
			defer wg.Done()
			с <- fmt.Sprintf("Goroutine %s", strconv.Itoa(i))
		}(c, i, &wg)
	}

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		}
	}

	wg.Wait()
	close(c)
}
