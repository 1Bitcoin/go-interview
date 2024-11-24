package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func ai() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}

var temp int
var mu = sync.RWMutex{}

func update() {
	value := ai()
	mu.Lock()
	temp = value
	mu.Unlock()
}

func t5() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	update()

	go func() {

	}()

	go func(ctx context.Context) {

		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(1 * time.Second):

			}

			update()
		}
	}(ctx)

	http.HandleFunc("/weather", func(writer http.ResponseWriter, request *http.Request) {
		mu.RLock()
		fmt.Println("temperature:", temp)
		mu.RUnlock()

	})

	http.HandleFunc("/stop2", func(writer http.ResponseWriter, request *http.Request) {
		cancel()
	})

	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}
