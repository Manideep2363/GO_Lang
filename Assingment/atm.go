package main

import "fmt"

func main() {
	var amount int
	fmt.Print("Enter the amount to withdraw: ")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Invalid amount. Please enter a positive value.")
		return
	}

	if amount%100 != 0 {
		fmt.Println("Invalid amount. Please enter a multiple of 100.")
		return
	}
	count := int(amount / 100)
	fmt.Printf("You will receive %d notes of 100.\n", count)

}
