package utils

import (
	"fmt"
	"time"
)

const MESSAGE = "! Seja bem-vindo de volta ao Gophimation"

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
