package main

import (
	"fmt"
	"sync"
	"time"
)

// throttle создает обертку над функцией f так, чтобы она вызывалась не чаще одного раза за duration миллисекунд.
func throttle(f func(), duration time.Duration) func() {
	var mu sync.Mutex
	var lastCalled time.Time

	return func() {
		mu.Lock()
		defer mu.Unlock()

		now := time.Now()
		if now.Sub(lastCalled) >= duration {
			lastCalled = now
			f()
		}
	}
}

func main() {
	// Пример функции, которую мы будем вызывать
	printHello := func() {
		fmt.Println("Hello!")
	}

	// Создаем "тормоз" для вызова printHello не чаще, чем раз в 1000 мс (1 секунда)
	throttledPrintHello := throttle(printHello, 1000*time.Millisecond)

	// Вызываем функцию несколько раз с интервалом больше 1000 мс
	for i := 0; i < 10; i++ {
		throttledPrintHello()
		time.Sleep(1200 * time.Millisecond) // Ждем 1.2 секунды между вызовами
	}
}
