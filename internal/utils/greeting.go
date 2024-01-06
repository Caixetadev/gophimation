package utils

import (
	"fmt"
	"time"
)

const MESSAGE = "! Seja bem-vindo de volta ao Gophimation"

// Greeting is a function that takes a user name and outputs a greeting message based on the time of day.
// It gets the current hour using the time package and uses a switch statement to determine the appropriate message.
// The function then prints the greeting message to the console.
func Greeting(userName string) {
	var message string

	switch hour := time.Now().Hour(); {
	case hour < 6:
		message = "Boa madrugada, " + userName + MESSAGE
	case hour < 12:
		message = "Bom dia, " + userName + MESSAGE
	case hour < 17:
		message = "Boa tarde, " + userName + MESSAGE
	default:
		message = "Boa noite, " + userName + MESSAGE
	}

	fmt.Println(message)
}
