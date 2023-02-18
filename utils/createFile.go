package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type UserConfig struct {
	Name string `json:"name"`
}

func CreateFile(fileName string) {
	var name string

	fmt.Println("Qual o seu nome?")

	fmt.Scanln(&name)

	file, err := os.Create(fileName)

	data := UserConfig{
		Name: name,
	}

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	user, _ := json.Marshal(data)

	_ = os.WriteFile(fileName, user, 0644)

	if err != nil {
		log.Fatal(err)
	}

	Clear()
}
