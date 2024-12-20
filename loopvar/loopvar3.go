package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
}
