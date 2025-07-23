package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter the name: ")
	fmt.Scanln(&name)
	fmt.Println("Vanakam " + name)
}
