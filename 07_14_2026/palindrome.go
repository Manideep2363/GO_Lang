package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	original := num
	rev := 0

	for num > 0 {
		d := num % 10
		rev = rev*10 + d
		num /= 10
	}
	if original == rev {
		fmt.Println(original, "is a palindrome")
	} else {
		fmt.Println(original, "is not a palindrome")
	}
}
