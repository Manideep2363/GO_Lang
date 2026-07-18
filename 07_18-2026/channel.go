package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 25
	}()

	x := <-ch

	fmt.Println(x)
}
