package main

import (
	"fmt"
	"time"
)

func number() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func alphabets() {
	for ch := 'A'; ch <= 'E'; ch++ {
		fmt.Printf("%c\n", ch)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go number()

	go alphabets()
	time.Sleep(300 * time.Millisecond)
}
