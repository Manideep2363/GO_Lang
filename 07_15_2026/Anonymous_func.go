package main

import "fmt"

func calculate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	add := func(x, y int) int {
		return x + y
	}
	subtract := func(x, y int) int {
		return x - y
	}
	multiply := func(x, y int) int {
		return x * y
	}
	fmt.Println(calculate(10, 5, add))
	fmt.Println(calculate(3, 5, subtract))
	fmt.Println(calculate(4, 6, multiply))
}
