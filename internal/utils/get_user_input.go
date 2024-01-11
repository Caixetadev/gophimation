package utils

import "fmt"

// GetUserInput retrieves user input by displaying a message in the console.
func GetUserInput(message string) int {
	var selectedOption int

	fmt.Println(message)
	fmt.Scanln(&selectedOption)

	return selectedOption
}
