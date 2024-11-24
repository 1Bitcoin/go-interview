package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Есть функция, работающая неопределенно долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(2000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

// Нужно изменить функцию обертку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.
const timeOut = 1 * time.Second

func predictableFunc(ctx context.Context) (int64, error) {
	start := time.Now()

	defer func() {
		fmt.Println("Execution time:", time.Since(start))
	}()

	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	ch := make(chan int64, 1)

	go func() {
		ch <- unpredictableFunc()
	}()

	select {
	case res := <-ch:
		return res, nil

	case <-ctx.Done():
		return 0, ctx.Err()
	}

	//select {
	//case res := <-ch:
	//	since := time.Since(start)
	//	fmt.Printf("time execution: %v\n", since)
	//	return res, nil
	//
	//case <-time.After(time.Duration(timeOut) * time.Second):
	//	since := time.Since(start)
	//	fmt.Printf("time execution: %v\n", since)
	//
	//	return 0, fmt.Errorf("timeout")
	//}
}

func t2() {
	fmt.Println("started")

	result, err := predictableFunc(context.Background())

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
