package utils

import (
	"log"
)

func OptionIsValid(animes []AnimeInfo, option int) {
	if option > len(animes) || option < 1 {
		Clear()
		log.Fatalf("Digite um número válido")
	}
}
