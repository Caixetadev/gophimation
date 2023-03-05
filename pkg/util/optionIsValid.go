package util

import (
	"log"
)

// OptionIsValid checks if the option selected by the user is valid for a given list of anime.
func OptionIsValid(anime []AnimeInfo, option int) {
	// Checks whether the selected option is invalid.
	if option > len(anime) || option < 1 {
		// Clears the screen and exits the program with an error message.
		Clear()
		log.Fatalf("Digite um número válido")
	}
}
