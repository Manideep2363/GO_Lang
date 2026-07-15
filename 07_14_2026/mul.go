package main

import "fmt"

func main() {
	var firstName, lastName string

	fmt.Print("Enter your first and last name: ")
	fmt.Scan(&firstName, &lastName)
	fmt.Println("Your name is:", firstName, lastName)
}
