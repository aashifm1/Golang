
package main

import (
	"fmt"
	"os"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Simple Calculator using Golang")
	
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not possible...")
			os.Exit(1)
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator...")
		os.Exit(1)
	}

	fmt.Printf("Result: %.2f\n", result)
}
