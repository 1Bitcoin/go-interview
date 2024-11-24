package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
)

func main() {
	targets := []string{
		"https://premier.ru",
		"https://premier.com",
		"http://premier.ru",
		"http://yandex.com",
		"htp://premier.ru",
		"http://premier.ru",
		"https://ya.ru",
		"https://twitter.com",
		"https://instagram.com",
	}

	fmt.Println(makeRequests(targets, 3))
}

func makeRequests(urls []string, n int) (succeed, failed int64) {
	wg := &sync.WaitGroup{}
	in := make(chan string)

	go func() {
		wg.Add(1)
		defer wg.Done()

		for _, target := range urls {
			in <- target
		}

		close(in)
	}()

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			//time.Sleep(time.Second)

			for item := range in {
				resp, err := http.Get(item)
				if err != nil {
					atomic.AddInt64(&failed, 1)
					break
				}

				if resp.StatusCode != 200 {
					atomic.AddInt64(&failed, 1)
					break
				}

				atomic.AddInt64(&succeed, 1)
			}
		}()
	}

	wg.Wait()

	return
}
