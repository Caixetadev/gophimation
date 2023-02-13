package search

import (
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

type AnimeInfo struct {
	Index int
	Name  string
	ID    string
}

func Search() string {
	c := colly.NewCollector(
		colly.AllowedDomains("www.anitube.site", "rr1---sn-gx5auxaxjvhxpgxap-btoe.googlevideo.com", "www.blogger.com"),
	)
	fname := os.Args[1]
	var option int
	var animes []AnimeInfo
	var animeSelected string

	URL := "https://www.anitube.site/?s=" + fname

	c.OnHTML(".aniItem", func(e *colly.HTMLElement) {
		href := e.ChildAttr("a", "href")
		name := e.ChildText(".aniItemNome")

		animes = append(animes, AnimeInfo{Name: name, ID: href, Index: e.Index})

		fmt.Printf("[%d] - nome do anime: %v. ID do anime: %v\n", e.Index, name, href)
	})

	c.Visit(URL)

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&option)

	for i, ai := range animes {
		if i == option {
			animeSelected = ai.ID
		}
	}

	return animeSelected
}
