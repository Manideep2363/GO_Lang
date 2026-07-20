package main

import (
	"fmt"
	"sync"
	"time"
)

var number = 100
var mu sync.RWMutex

func reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.RLock()
	defer mu.RUnLock()

	fmt.Println("reader", id, "reads: ", number)

	time.Sleep(2 * time.Second)

}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go reader(i, &wg)
	}
	wg.Wait()
}
