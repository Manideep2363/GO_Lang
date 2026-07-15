package main

import "fmt"

func s(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	fmt.Println(s(1, 2, 3, 4, 5))
}
