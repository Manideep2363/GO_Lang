package main

import "fmt"

func main() {
	var num [5]int
	for i := 0; i < len(num); i++ {
		fmt.Print("Enter value for index ", i, ": ")
		fmt.Scan(&num[i])
	}
	for i := 0; i < len(num); i++ {
		fmt.Print("Value at index ", i, " is: ", num[i], "\n")
	}

	largest := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] > largest {
			largest = num[i]
		}
	}
	fmt.Println("Largest number is:", largest)

	smallest := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] < smallest {
			smallest = num[i]
		}
	}
	fmt.Println("Smallest number is:", smallest)
}
