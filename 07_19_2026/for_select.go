package main

import (
	"fmt"
	"time"
)

func main() {
	email := make(chan string)
	sms := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		email <- fmt.Sprintf("email 1")
	}()

	go func() {
		time.Sleep(2 * time.Second)
		sms <- fmt.Sprintf("Sms 1")
	}()
}
