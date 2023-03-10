package search

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Caixetadev/gophimation/pkg/configs"
	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/Caixetadev/gophimation/pkg/models"
	"github.com/Caixetadev/gophimation/pkg/util"
	"github.com/gocolly/colly"
)

// Search does the search for the anime
func Search() string {
	searchTerm := os.Args
	c := configs.Colly()

	var URL string

	for i := 1; i < len(searchTerm); i++ {
		URL += fmt.Sprintf("%spesquisa?titulo=%s&searchTerm=%s", constants.URL_BASE, searchTerm[i]+"+", searchTerm[i]+"+")
	}

	var selectedOption int

	var anime []models.Anime

	c.OnHTML(".list-animes article", func(h *colly.HTMLElement) {
		episode := h.ChildAttr("a", "title")
		urlAnime := h.ChildAttr("a", "href")

		anime = append(anime, models.Anime{Name: episode, URL: strings.TrimPrefix(urlAnime, constants.URL_BASE)})
	})

	c.Visit(URL)

	for i, item := range anime {
		fmt.Printf("[%02d] - %v\n", i+1, item.Name)
	}

	fmt.Println("\ncoloque um numero para assistir")

	if len(anime) == 0 {
		log.Fatal("Não foi possivel achar o anime")
	}

	fmt.Scanln(&selectedOption)

	util.OptionIsValid(anime, selectedOption)

	fmt.Println()

	util.Clear()

	return anime[selectedOption-1].URL
}
