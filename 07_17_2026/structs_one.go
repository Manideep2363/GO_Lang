package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Marks int
}

func main() {

	// var students []Student
	var Name string
	var Age int
	var Marks int
	// var s Student
	var i int
	for i = 0; i < 3; i++ {
		fmt.Println(i)
	}

	// for _, s := range students {
	// 	fmt.Println("Name: ", s.Name, ", Age: ", s.Age, ", Marks: ", s.Marks)
	// }
}
