package main

import (
	"fmt"
)

func display[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}

}

func main() {
	display([]int{10, 20, 30})
	display([]string{"Apple", "Banana", "Mango"})
}
