package utils

import (
	"log"

	"github.com/Caixetadev/gophimation/internal/entity"
)

// OptionIsValid checks if the option selected by the user is valid for a given list of anime.
func OptionIsValid(anime []entity.Anime, selectedOption int) {
	// Checks whether the selected option is invalid.
	if selectedOption <= 0 || selectedOption > len(anime) {
		// Clears the screen and exits the program with an error message.
		log.Fatalf("Digite um número válido")
	}
}
