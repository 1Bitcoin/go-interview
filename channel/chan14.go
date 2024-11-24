package main

import (
	"fmt"
	"sync"
)

type (
	Broker struct {
		subscribers []chan string
		mu          sync.RWMutex
		closed      bool
	}
)

func NewBroker() *Broker {
	return &Broker{
		mu: sync.RWMutex{},
	}
}

func (s *Broker) Subscribe() <-chan string {
	ch := make(chan string)
	s.subscribers = append(s.subscribers, ch)

	return ch
}

func (s *Broker) Publish(value string) {
	for _, ch := range s.subscribers {
		ch <- value
	}
}

func (s *Broker) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return
	}

	for _, ch := range s.subscribers {
		close(ch)
	}

	s.closed = true
}

// Результатом работы программы должен быть вывод строк:
// event #1
// event #2
func main() {
	broker := NewBroker()

	defer broker.Close()

	subscriber := broker.Subscribe()

	go func() {
		for event := range subscriber {
			fmt.Println(event)
		}
	}()

	broker.Publish("event #1")
	broker.Publish("event #2")
}
