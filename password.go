package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func gen_psw(length int, includeUpper bool, includeDigits bool, includeSpecial bool) string {
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	specialChars := "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"

	charSet := lowercase

	if includeUpper {
		charSet += uppercase
	}

	if includeDigits {
		charSet += digits
	}

	if includeSpecial {
		charSet += specialChars
	}

	rand.Seed(time.Now().UnixNano()) // Generate random
	password := strings.Builder{}
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charSet))
		password.WriteByte(charSet[randomIndex])
	}
	return password.String()
}

func main() {
	var length int
	var includeUpper, includeDigits, includeSpecial bool

	fmt.Println("Enter the password length: ")
	fmt.Scanln(&length)

	fmt.Print("Include uppercase letters? (y/n): ")
	var upperChoice string
	fmt.Scanln(&upperChoice)
	includeUpper = upperChoice == "y" || upperChoice == "Y"

	fmt.Print("Include digits? (y/n): ")
	var digitsChoice string
	fmt.Scanln(&digitsChoice)
	includeDigits = digitsChoice == "y" || digitsChoice == "Y"

	fmt.Print("Include special characters? (y/n): ")
	var specialChoice string
	fmt.Scanln(&specialChoice)
	includeSpecial = specialChoice == "y" || specialChoice == "Y"

	password := gen_psw(length, includeUpper, includeDigits, includeSpecial)
	fmt.Printf("Generated Password: %s\n", password)
}
