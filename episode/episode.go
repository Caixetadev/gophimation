package episode

import (
	"fmt"

	"github.com/Caixetadev/gophimation/config"
	"github.com/Caixetadev/gophimation/search"
	"github.com/gocolly/colly/v2"
)

type EpisodeInfo struct {
	Index int
	Name  string
	ID    string
}

func SelectEpisode() string {
	c := config.Colly()

	var episodes []EpisodeInfo
	var option int
	var episodeSelected string

	c.OnHTML(".pagAniListaContainer a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		name := e.Attr("title")

		episodes = append(episodes, EpisodeInfo{Name: name, ID: href, Index: e.Index})

		fmt.Printf("[%d] -  %v\n", e.Index+1, name)
	})

	URL := search.Search()

	c.Visit(URL)

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&option)

	for i, ai := range episodes {
		if (i + 1) == option {
			episodeSelected = ai.ID
		}
	}

	return episodeSelected
}
