package main

import (
	"context"

	"fmt"

	"sync"

	"time"
)

const rubToUsd = 100

type Snippet struct {
	name string

	price float64
}

func NewSnippet(name string, price float64) Snippet {

	return Snippet{

		name: name,

		price: price,
	}

}

func getDescriptionIntegration(ctx context.Context, itemId int) string {

	time.Sleep(2 * time.Second)

	return "good snippet"

}

func getPriceIntegration(ctx context.Context, itemId int) float64 {

	time.Sleep(2 * time.Second)

	return 228

}

func convertDescription(description string) string {

	return description + "beutify"

}

func convertPrice(price float64) float64 {

	return price * rubToUsd

}

var wg sync.WaitGroup

func snippet(itemId int) Snippet {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	defer cancel()

	wg.Add(2)

	var price float64

	var desc string

	go func() {

	}()

	go func(ctx context.Context) {

		defer wg.Done()

		select {

		case <-ctx.Done():

			fmt.Println("timeout")

			return

		case <-time.After(1 * time.Second):

			fmt.Println("timeout")

			return

		default:

			desc = getDescriptionIntegration(ctx, itemId)

			desc = convertDescription(desc)

		}

	}(ctx)

	go func(ctx context.Context) {

		defer wg.Done()

		select {

		case <-ctx.Done():

			fmt.Println("timeout")

			return

		default:

			price = getPriceIntegration(ctx, itemId)

			price = convertPrice(price)

		}

	}(ctx)

	wg.Wait()

	return NewSnippet(desc, price)

}
