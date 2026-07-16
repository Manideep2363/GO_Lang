package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func cal(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	result := cal(3, 4, add)
	fmt.Println(result)
}
