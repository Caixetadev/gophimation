package mostwatched

import (
	"fmt"
	"log"

	"github.com/Caixetadev/gophimation/config"
	"github.com/Caixetadev/gophimation/utils"
	"github.com/gocolly/colly/v2"
)

// MostWatched prints the most viewed anime of the week and returns its url
func MostWatched() string {
	c := config.Colly()

	var animes []utils.AnimeInfo
	var option int
	var animeSelected string

	c.OnHTML(".owl-carousel-semana .containerAnimes", func(h *colly.HTMLElement) {
		animes = utils.ScrapeAnimeInfo(h)
	})

	if err := c.Visit("https://animefire.net"); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&option)

	utils.OptionIsValid(animes, option)

	for index, anime := range animes {
		if (index + 1) == option {
			animeSelected = anime.ID
			break
		}
	}

	utils.Clear()

	return animeSelected
}
