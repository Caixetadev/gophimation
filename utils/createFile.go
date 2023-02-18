package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Caixetadev/gophimation/constants"
)

type UserConfig struct {
	Name string `json:"name"`
}

func CreateFile() {
	var name string

	fmt.Println("Qual o seu nome?")

	fmt.Scanln(&name)

	file, err := os.Create(constants.FILE_NAME)

	data := UserConfig{
		Name: name,
	}

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	user, _ := json.Marshal(data)

	_ = os.WriteFile(constants.FILE_NAME, user, 0644)

	if err != nil {
		log.Fatal(err)
	}

	Clear()
}
