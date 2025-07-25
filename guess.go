package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Generate random
	target := rand.Intn(10) + 1      // Setting the limit (1-10)
	var guess int

	fmt.Println("Guess the number Game!!")
	fmt.Println("Guess the number from 1 to 10...")

	for {
		fmt.Println("Enter your guess:")
		fmt.Scanln(&guess)

		if guess < target {
			fmt.Println("Too low! Try again...")
		} else if guess > target {
			fmt.Println("Too high! Try again...")
		} else {
			fmt.Println("Congratulation! You guessed the correct number!!")
			break
		}
	}
}
