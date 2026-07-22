package main

import (
	"fmt"
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	n := len(s.items)

	item := s.items[n-1]
	s.items = s.items[:n-1]

	return item
}

func main() {
	var s Stack[int]

	s.Push(10)

	s.Push(20)

	s.Push(30)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}
