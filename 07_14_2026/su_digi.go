package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	original := num
	s := 0

	for num > 0 {
		d := num % 10
		s += d
		num /= 10
	}
	fmt.Println("Sum of digits of", original, "is", s)
}
