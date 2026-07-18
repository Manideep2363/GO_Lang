package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 100
	ch <- 200
	ch <- 300

	fmt.Println(len(ch))
	fmt.Println(cap(ch))

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
