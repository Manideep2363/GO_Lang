package main

import "fmt"

func cal(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	s, d := cal(10, 5)
	fmt.Println("Sum:", s)
	fmt.Println("Difference:", d)
}
