package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter any number: ")
	fmt.Scan(&num)

	var operator string
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	var num2 int
	fmt.Print("Enter another number: ")
	fmt.Scan(&num2)

	switch {
	case operator == "+":
		fmt.Print("Addition of two numbers is: ", num+num2)
	case operator == "-":
		fmt.Print("Subtraction of two numbers is: ", num-num2)
	case operator == "*":
		fmt.Print("Multiplication of two numbers is: ", num*num2)
	case operator == "/":
		if num2 == 0 {
			fmt.Print("Division by zero is not allowed.")
		} else {
			fmt.Print("Division of two numbers is: ", num/num2)
		}
	default:
		fmt.Print("Invalid operator.")

	}

}
