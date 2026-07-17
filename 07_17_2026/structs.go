package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Name  string
	Age   int
	Marks int
}

func main() {

	// var students []Student
	for i := 0; i < 3; i++ {
		// var s Student

		fmt.Println("Enter name, age and marks of student separated by comma: ")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			fmt.Println("Invalid input! Please enter in the format: Name,Age,Marks")
			i-- // Retry the same student
			continue
		}
		name := strings.TrimSpace(parts[0])
		age := 0
		marks := 0
		fmt.Sscanf(parts[1], "%d", &age)
		fmt.Sscanf(parts[2], "%d", &marks)

		fmt.Println(name)

	}
}
