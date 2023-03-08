package search

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Caixetadev/gophimation/pkg/models"
	mostWatched "github.com/Caixetadev/gophimation/pkg/mostWatched"
	"github.com/Caixetadev/gophimation/pkg/util"
)

// Search does the search for the anime
func Search() string {
	flags := os.Args

	var hasArgs = len(flags) == 1

	if hasArgs {
		URL := mostWatched.MostWatched()

		return URL
	}

	URL := "http://localhost:8000/search/"

	for i := 1; i < len(flags); i++ {
		URL += flags[i] + "-"
	}

	var option int

	resp, err := http.Get(URL)

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
		fmt.Printf("[%d] - %v\n", i+1, item.Name)
	}

	fmt.Println("\ncoloque um numero para assistir")

	if len(anime) == 0 {
		util.Clear()
		log.Fatal("Não foi possivel achar o anime")
	}

	fmt.Scanln(&option)

	util.OptionIsValid(anime, option)

	fmt.Println()

	return anime[option-1].URL
}
