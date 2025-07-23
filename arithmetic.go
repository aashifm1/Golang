package main

import "fmt"

func main() {
	var num1, num2 float64

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	sum := num1 + num2
	difference := num1 - num2
	product := num1 * num2

	fmt.Println(sum)

	fmt.Println(difference)

	fmt.Println(product)

}
