package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i == 20 {
			break
		}

		if i%3 == 0 {
			continue
		}

		fmt.Println(i)

	}
}
