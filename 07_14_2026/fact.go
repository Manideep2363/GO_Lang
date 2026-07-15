package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	var s = 1
	for i := 1; i <= num; i++ {
		s *= i
	}
	fmt.Println("Factorial of", num, "is", s)
}
