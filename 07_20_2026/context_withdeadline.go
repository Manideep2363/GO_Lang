package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Dead Line reached")
			return
		default:
			fmt.Println("working")
			time.Sleep(time.Second)
		}
	}

}

func main() {
	deadline := time.Now().Add(3 * time.Second)

	ctx, cancel := context.WithDeadline(
		context.Background(),
		deadline,
	)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go worker(ctx, &wg)

	wg.Wait()
}
