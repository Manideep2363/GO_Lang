package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int)

	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for j := 1; j <= 6; j++ {
		jobs <- j
	}

	close(jobs)
	wg.Wait()
}
