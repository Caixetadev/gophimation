package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// UserConfig stores user configuration information, including their name.
type UserConfig struct {
	Name string `json:"name"`
}

// CreateFile prompts the user to enter their name, creates a new file
// with the specified name, and writes the user's name to the file as
// JSON data in a UserConfig struct format.
func CreateFile(rootDir string, filename string) {
	var name string

	fmt.Println("Qual o seu nome?")

	fmt.Scanln(&name)

	err := os.Mkdir(rootDir, 0777)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(rootDir + filename)

	data := UserConfig{
		Name: name,
	}

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	user, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(rootDir+filename, user, 0644)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	Clear()
}
