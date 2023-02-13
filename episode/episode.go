package episode

import (
	"fmt"

	"github.com/Caixetadev/ani-go/search"
	"github.com/gocolly/colly/v2"
)

type EpisodeInfo struct {
	Index int
	Name  string
	ID    string
}

func SelectEpisode() string {
	c := colly.NewCollector(
		colly.AllowedDomains("www.anitube.site", "rr1---sn-gx5auxaxjvhxpgxap-btoe.googlevideo.com", "www.blogger.com"),
	)

	var episodes []EpisodeInfo
	var option int
	var episodeSelected string

	c.OnHTML(".pagAniListaContainer a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		name := e.Attr("title")

		episodes = append(episodes, EpisodeInfo{Name: name, ID: href, Index: e.Index})

		fmt.Printf("[%d] - nome do anime: %v. ID do anime: %v\n", e.Index, name, href)
	})

	URL := search.Search()

	c.Visit(URL)

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&option)

	for i, ai := range episodes {
		if i == option {
			episodeSelected = ai.ID
		}
	}

	return episodeSelected
}
