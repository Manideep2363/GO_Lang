package main

import "fmt"

func Producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func Consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println((value))
	}
}

func main() {
	ch := make(chan int)

	go Producer(ch)
	Consumer(ch)

}
