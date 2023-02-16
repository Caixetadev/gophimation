package episode

import (
	"fmt"

	"github.com/Caixetadev/gophimation/config"
	"github.com/Caixetadev/gophimation/search"
	"github.com/Caixetadev/gophimation/utils"
	"github.com/gocolly/colly/v2"
)

func SelectEpisode() string {
	c := config.Colly()

	var episodes []utils.AnimeInfo
	var option int
	var episodeSelected string

	c.OnHTML(".pagAniListaContainer a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		name := e.Attr("title")

		episodes = append(episodes, utils.AnimeInfo{Name: name, ID: href, Index: e.Index})

		fmt.Printf("[%d] -  %v\n", e.Index+1, name)
	})

	URL := search.Search()

	c.Visit(URL)

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&option)

	utils.OptionIsValid(episodes, option)

	for i, ai := range episodes {
		if (i + 1) == option {
			episodeSelected = ai.ID
		}
	}

	return episodeSelected
}
