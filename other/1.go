package main

import (
	"fmt"
	"sync"
	"time"
)

const numRequests = 10000

var count int

var waitGroup sync.WaitGroup
var mutex sync.Mutex

func networkRequest() {
	time.Sleep(time.Millisecond) // Эмуляция сетевого запроса.

	mutex.Lock()
	count++
	mutex.Unlock()
}

func t1() {

	start := time.Now()

	waitGroup.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer waitGroup.Done()
			networkRequest()
		}()
	}

	waitGroup.Wait()

	since := time.Since(start)
	fmt.Println(count, " ", since)
}
