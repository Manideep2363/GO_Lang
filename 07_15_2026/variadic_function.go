package main

import "fmt"

func display(numbers ...int) {
	fmt.Println(numbers)
}

func main() {
	display(10)
	display(10, 20)
	display(10, 20, 30)
}
