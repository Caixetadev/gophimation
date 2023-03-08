package episode

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Caixetadev/gophimation/pkg/models"
	"github.com/Caixetadev/gophimation/pkg/util"
)

func SelectEpisode(URL string) string {
	var option int

	resp, err := http.Get("http://localhost:8000/episodes/" + URL)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var anime []models.Anime
	err = json.NewDecoder(resp.Body).Decode(&anime)

	if err != nil {
		log.Fatal(err)
	}

	for i, item := range anime {
		fmt.Printf("[%02d] - %v\n", i+1, item.Name)
	}

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&option)

	util.OptionIsValid(anime, option)

	util.Clear()

	fmt.Println("Carregando...")

	return anime[option-1].URL
}
