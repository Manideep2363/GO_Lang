package main

import "fmt"

func Producer(ch chan<- int) {
	ch <- 100
}

func Consumer(ch <-chan int) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan int)

	go Producer(ch)

	Consumer(ch)

}
