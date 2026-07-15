package main

import "fmt"

func main() {
	var numbers [5]int

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Enter value for index %d: ", i)
		fmt.Scan(&numbers[i])
	}

	fmt.Println("Array:", numbers)
}
