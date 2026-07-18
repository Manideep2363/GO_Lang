package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker", id, "Started")
	time.Sleep(2 * time.Second)
	fmt.Println("Worker", id, "finished")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("start")
	fmt.Println("all workers finised")
	fmt.Println("end")
}
