package mostwatched

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Caixetadev/gophimation/pkg/models"
	"github.com/Caixetadev/gophimation/pkg/util"
)

// MostWatched prints the most viewed anime of the week and returns its url
func MostWatched() string {
	var option int

	resp, err := http.Get("http://localhost:8000/most-watched")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var anime []models.AnimeInfo
	err = json.NewDecoder(resp.Body).Decode(&anime)

	if err != nil {
		log.Fatal(err)
	}

	for i, item := range anime {
		fmt.Printf("[%d] - %v\n", i+1, item.Name)
	}

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&option)

	util.OptionIsValid(anime, option)

	return anime[option-1].ID
}
