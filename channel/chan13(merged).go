package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		close(ch1)
	}()

	go func() {
		ch1 <- 2
		close(ch2)
	}()

	for val := range mergeChannels(ch1, ch2) {
		fmt.Println(val)
	}
}

func mergeChannels(channels ...<-chan int) chan int {

	merged := make(chan int)

	wg := sync.WaitGroup{}

	// Определяем функцию для чтения значений из канала
	// и отправки в объединенный канал
	read := func(c <-chan int) {
		defer wg.Done()

		for val := range c {
			merged <- val
		}
	}

	// Добавляем каждый канал в группу ожидания
	wg.Add(len(channels))

	// Запускаем чтение значений из каждого канала в отдельной горутине
	for _, ch := range channels {
		go read(ch)
	}

	// Запускаем горутину, которая закрывает объединенный канал
	// после того, как все горутины закончат свою работу
	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
