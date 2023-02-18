package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ReadFile(fileName string) {
	data, err := os.ReadFile(fileName)

	var user UserConfig

	if errUnmarshal := json.Unmarshal(data, &user); errUnmarshal != nil {
		log.Fatalln(errUnmarshal)
	}

	Greeting(user.Name)

	fmt.Println()

	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
}
