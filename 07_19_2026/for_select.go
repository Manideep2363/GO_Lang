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
		email <- "Email 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		sms <- "sms 1"
	}()

	received := 0

	for {
		select {
		case msg := <-email:
			fmt.Println(msg)
			received++

		case msg := <-sms:
			fmt.Println(msg)
			received++

		}

		if received == 2 {
			fmt.Print("done")
			break
		}

	}

}
