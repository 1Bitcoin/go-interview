package main

import (
	"fmt"
	"sync"
)

type Pool struct {
}

func NewPool() *Pool {
	return &Pool{}
}

func (p *Pool) work(id int, f func(int) int, jobs <-chan int, results chan<- int, group *sync.WaitGroup) {
	for j := range jobs {
		results <- f(j)
	}

	group.Done()
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	multiplier := func(x int) int {
		return x * 10
	}

	pool := NewPool()

	go func() {
		wg := sync.WaitGroup{}

		wg.Add(3)
		for w := 1; w <= 3; w++ {
			go pool.work(w, multiplier, jobs, results, &wg)
		}
		wg.Wait()
		close(results)
	}()

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for res := range results {
		fmt.Println(res)
	}
}
