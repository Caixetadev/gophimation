package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// ReadFile reads the specified file and unmarshals its JSON content
// into a UserConfig struct. It then calls the Greeting function to
// display a welcome message for the user.
func ReadFile(rootDir string, fileName string) {
	fmt.Println(rootDir + fileName)
	data, err := os.ReadFile(rootDir + fileName)

	var user UserConfig

	if errUnmarshal := json.Unmarshal(data, &user); errUnmarshal != nil {
		log.Fatalln(errUnmarshal)
	}

	Clear()

	Greeting(user.Name)

	fmt.Print("\n")

	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
}
