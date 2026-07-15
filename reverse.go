package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	original := n
	rev := 0

	for n > 0 {
		d := n % 10
		rev = rev*10 + d
		n /= 10
	}
	fmt.Println("Reverse of", original, "is", rev)
}
