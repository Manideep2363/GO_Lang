package main

import (
	"fmt"
)

type Box[T any] struct {
	value T
}

func (b *Box[T]) Set(value T) {
	b.value = value
}

func (b *Box[T]) Get() T {

	return b.value
}

func main() {
	var b Box[int]

	b.Set(100)
	fmt.Println(b.Get())

	var s Box[string]

	s.Set("go")

	fmt.Println(s.Get())

}
