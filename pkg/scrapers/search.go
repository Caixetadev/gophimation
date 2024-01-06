package scrapers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Caixetadev/gophimation/config"
	"github.com/Caixetadev/gophimation/internal/entity"
	"github.com/Caixetadev/gophimation/internal/utils"
	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/gocolly/colly"
)

// Search does the search for the anime
func Search() string {
	c := config.Colly()

	searchTerm := strings.Join(os.Args[1:], "+")

	URL := fmt.Sprintf("%spesquisa?titulo=%s&searchTerm=%s", constants.URL_BASE, searchTerm, searchTerm)

	var selectedOption int

	var anime []entity.Anime

	c.OnHTML(".list-animes article", func(h *colly.HTMLElement) {
		fmt.Printf("[%02d] - %s\n", h.Index+1, h.ChildAttr("a", "title"))

		anime = append(anime, entity.Anime{URL: strings.TrimPrefix(h.ChildAttr("a", "href"), constants.URL_BASE)})
	})

	if err := c.Visit(URL); err != nil {
		log.Fatal(err)
	}

	if len(anime) == 0 {
		log.Fatal("Não foi possivel achar o anime")
	}

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&selectedOption)

	utils.OptionIsValid(anime, selectedOption)

	utils.Clear()

	return anime[selectedOption-1].URL
}
