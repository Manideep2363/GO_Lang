package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(2 * time.Second)

	for i := 1; i <= 3; i++ {
		<-timer.C
		fmt.Println("Tick:", i)

	}
	timer.Stop()
	fmt.Println("Ticker Stopped")
}
