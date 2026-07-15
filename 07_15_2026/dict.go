package main

import "fmt"

func main() {
	var s string
	fmt.Print("Enter a string: ")
	fmt.Scan(&s)
	d := make(map[rune]int)
	for _, c := range s {
		d[c]++
	}
	for char, count := range d {
		fmt.Printf("Character '%c' occurs %d times\n", char, count)
	}
}
