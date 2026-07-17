package main

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	Name  string
	Age   int
	Marks int
}

func main() {
	var students []Student
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 3; i++ {
		var s Student

		fmt.Printf("Enter Student %d (Name,Age,Marks): ", i+1)

		line, _ := reader.ReadString('\n')

		_, err := fmt.Sscanf(line, "%[^,],%d,%d", &s.Name, &s.Age, &s.Marks)
		if err != nil {
			fmt.Println("Invalid input! Please enter in the format: Name,Age,Marks")
			i-- // Retry the same student
			continue
		}

		students = append(students, s)
	}

	fmt.Println("\nStudent Details:")
	for _, s := range students {
		fmt.Printf("%+v\n", s)
	}
}
