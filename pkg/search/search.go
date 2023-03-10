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
	c := configs.Colly()

	searchTerm := strings.Join(os.Args[1:], "+")

	URL := fmt.Sprintf("%spesquisa?titulo=%s&searchTerm=%s", constants.URL_BASE, searchTerm, searchTerm)

	var selectedOption int

	var anime []models.Anime

	c.OnHTML(".list-animes article", func(h *colly.HTMLElement) {
		fmt.Printf("[%02d] - %s\n", h.Index+1, h.ChildAttr("a", "title"))

		anime = append(anime, models.Anime{URL: strings.TrimPrefix(h.ChildAttr("a", "href"), constants.URL_BASE)})
	})

	c.Visit(URL)

	if len(anime) == 0 {
		log.Fatal("Não foi possivel achar o anime")
	}

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&selectedOption)

	util.OptionIsValid(anime, selectedOption)

	util.Clear()

	return anime[selectedOption-1].URL
}
