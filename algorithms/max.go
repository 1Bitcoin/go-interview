package main

import (
	"fmt"

	"sync"
)

func main() {

	var max int

	var wg sync.WaitGroup

	var mutex sync.Mutex

	wg.Add(10000)

	for i := 10000; i > 0; i-- {

		go func(i int) {

			defer wg.Done()

			mutex.Lock()

			if i%2 == 0 && i > max {

				max = i

			}

			mutex.Unlock()

		}(i)

	}

	wg.Wait()

	fmt.Println("max:", max)

}
