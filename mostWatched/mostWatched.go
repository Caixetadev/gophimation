package mostwatched

import (
	"fmt"

	"github.com/Caixetadev/gophimation/config"
	"github.com/Caixetadev/gophimation/utils"
	"github.com/gocolly/colly/v2"
)

func MostWatched() string {
	c := config.Colly()

	var animes []utils.AnimeInfo
	var option int
	var animeSelected string

	c.OnHTML(".aniContainer .main-carousel .aniItem", func(h *colly.HTMLElement) {
		animes = utils.ScrapeAnimeInfo(h)
	})

	c.Visit("https://www.anitube.site/")

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
