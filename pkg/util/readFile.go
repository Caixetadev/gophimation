package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// ReadFile reads the specified file and unmarshals its JSON content
// into a UserConfig struct. It then calls the Greeting function to
// display a welcome message for the user.
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
