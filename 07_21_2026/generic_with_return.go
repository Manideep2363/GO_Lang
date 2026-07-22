package main

import "fmt"

type Number interface {
	~int | ~float64
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Min(15, 8))
	fmt.Println(Min(4.5, 9.2))
}
