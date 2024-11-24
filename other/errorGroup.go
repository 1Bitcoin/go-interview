package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

// Функция для выполнения HTTP-запроса
func fetchURL(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // Закрыть тело ответа после прочтения

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch %s: %s", url, resp.Status)
	}

	fmt.Printf("Successfully fetched: %s\n", url)
	return nil
}

func main() {
	var g errgroup.Group

	ctx := context.Background()

	// Список URL для получения
	urls := []string{
		"https://example.com",
		"htt://example.org",
		"https://example.net",
	}

	for _, url := range urls {
		// Создаем замыкание для захвата переменной url
		url := url // захватываем значение в замыкание
		g.Go(func() error {
			return fetchURL(ctx, url)
		})
	}

	// Ждем завершения всех горутин и проверяем наличие ошибок
	if err := g.Wait(); err != nil {
		fmt.Printf("Error fetching URLs: %v\n", err)
	} else {
		fmt.Println("All URLs fetched successfully.")
	}
}
