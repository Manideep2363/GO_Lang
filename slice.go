package main

import "fmt"

func main() {
	var num []int
	var n int
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&n)
	var s int
	for i := 0; i < n; i++ {
		fmt.Print("Enter value for index ", i, ": ")
		var value int
		fmt.Scan(&value)
		s += value
		num = append(num, value)
	}
	fmt.Println("The slice is:", num)
	fmt.Printf("Length of the slice is: %d\n", len(num))
	fmt.Printf("Sum of the elements is: %d\n", s)
}
