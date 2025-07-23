package main

import "fmt"

func main() {
	var num1, num2, result float64
	var operator string

	fmt.Println("Simpe Calculator using Golang")

	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Println("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	switch operator {

	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Division by zero is not possible...")
			return
		}

	default:
		fmt.Println("Invalifd operator")
		return

	}

	fmt.Printf("Result: %.2f\n", result)
}
