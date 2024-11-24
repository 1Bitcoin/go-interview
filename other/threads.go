package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func threads() {
	var wg sync.WaitGroup

	// Количество вызовов API, которые мы хотим выполнить
	numCalls := 2

	// Создаем канал, чтобы сохранить ответы API
	responses := make(chan *Post, numCalls)

	// Устанавливаем количество ожидаемых вызовов в группе ожидания
	wg.Add(numCalls)

	// Функция для выполнения вызова API
	apiCall := func(url string) {
		defer wg.Done()

		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}

		}(resp.Body)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		post := &Post{}
		err = json.Unmarshal(body, post)
		if err != nil {
			log.Fatal(err)
		}

		responses <- post
	}

	// Первый многопоточный вызов API
	go apiCall("https://jsonplaceholder.typicode.com/posts/1")

	// Второй многопоточный вызов API
	go apiCall("https://jsonplaceholder.typicode.com/posts/2")

	// Ожидаем завершения всех вызовов
	wg.Wait()

	// Закрываем канал после завершения всех вызовов
	close(responses)

	fmt.Printf("\n")
	for post := range responses {
		fmt.Printf("Получены данные: %+v\n", *post)
		fmt.Printf("\n")
	}
}
