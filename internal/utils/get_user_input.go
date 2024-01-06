package utils

import "fmt"

func GetUserInput(message string) int {
	var selectedOption int

	fmt.Println(message)
	fmt.Scanln(&selectedOption)

	return selectedOption
}
