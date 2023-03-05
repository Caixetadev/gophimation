package mostwatched

import (
	"fmt"
	"log"

	"github.com/Caixetadev/gophimation/configs"
	"github.com/Caixetadev/gophimation/pkg/util"
	"github.com/gocolly/colly/v2"
)

// MostWatched prints the most viewed anime of the week and returns its url
func MostWatched() string {
	c := configs.Colly()

	var anime []util.AnimeInfo
	var option int
	var animeSelected string

	c.OnHTML(".owl-carousel-semana .containerAnimes", func(h *colly.HTMLElement) {
		anime = util.ScrapeAnimeInfo(h)
	})

	if err := c.Visit("https://animefire.net"); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&option)

	util.OptionIsValid(anime, option)

	for index, anime := range anime {
		if (index + 1) == option {
			animeSelected = anime.ID
			break
		}
	}

	util.Clear()

	return animeSelected
}
