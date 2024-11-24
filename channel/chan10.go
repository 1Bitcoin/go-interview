package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			m[key] = key * key
		}(i)
	}

	wg.Wait()

	for key, val := range m {
		fmt.Printf("Key: %d, Value: %d\n", key, val)
	}
}
