package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic произошла. Сообщение:", r)
		}
	}()

	defer fmt.Println("Это сообщение будет выведено перед завершением программы")
	fmt.Println("Выполнение до паники")

	panic("Произошла паника")
}
