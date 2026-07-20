package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Program started")
	timer := time.NewTimer(2 * time.Second)

	<-timer.C

	fmt.Println("2 seconds Completed")
	fmt.Println("Program finised")
}
