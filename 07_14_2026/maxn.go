package main

import "fmt"

func main() {
	num := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(num); i++ {
		fmt.Print("enter value for index ", i, ": ")
		fmt.Scan(&num[i])
	}
	largest := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] > largest {
			largest = num[i]
		}
	}
	fmt.Println("Largest number is:", largest)
}
