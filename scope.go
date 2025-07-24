package main

import "fmt"

var global int = 1002 // Global Variable

func main() {
	var local int = 1003 // Local Variable
	fmt.Println("Global Variable - Inside function:", global)
	fmt.Println("Local Variable:", local)
	outside()
}

func outside() {
	fmt.Println("Global Variable - Outside function:", global)

}
