package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Negative number")
	} else if num == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Positive number")
	}

	if num != 0 {
		if num%2 == 0 {
			fmt.Println("Even number")
		} else {
			fmt.Println("Odd number")
		}
	}
}
