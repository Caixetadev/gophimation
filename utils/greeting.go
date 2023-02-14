package utils

import (
	"fmt"
	"time"
)

func Greeting(userName string) {
	var message string

	switch hour := time.Now().Hour(); {
	case hour < 6:
		message = "Boa madrugada, " + userName + "! Seja bem-vindo de volta ao ANI-GO"
	case hour < 12:
		message = "Bom dia, " + userName + "! Seja bem-vindo de volta ao ANI-GO"
	case hour < 17:
		message = "Boa tarde, " + userName + "! Seja bem-vindo de volta ao ANI-GO"
	default:
		message = "Boa noite, " + userName + "! Seja bem-vindo de volta ao ANI-GO"
	}

	fmt.Println(message)
}
