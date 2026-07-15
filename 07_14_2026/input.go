package main

import "fmt"

func main() {
	var age int

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println("Your age is:", age)
}
