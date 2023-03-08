package util

import (
	"log"

	"github.com/Caixetadev/gophimation/pkg/models"
)

// OptionIsValid checks if the option selected by the user is valid for a given list of anime.
func OptionIsValid(anime []models.Anime, option int) {
	// Checks whether the selected option is invalid.
	if option <= 0 || option > len(anime) {
		// Clears the screen and exits the program with an error message.
		Clear()
		log.Fatalf("Digite um número válido")
	}
}
