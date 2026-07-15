package main

import "fmt"

func main() {
	num := 1
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(num, " ")
			num++
		}
		fmt.Println()
	}
}
