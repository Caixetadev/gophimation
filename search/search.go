package search

import (
	"fmt"
	"os"

	"github.com/Caixetadev/gophimation/config"
	"github.com/Caixetadev/gophimation/utils"
	"github.com/gocolly/colly/v2"
)

type AnimeInfo struct {
	Index int
	Name  string
	ID    string
}

func Search() string {
	c := config.Colly()

	fname := os.Args
	var option int
	var animes []AnimeInfo
	var animeSelected string
	URL := "https://www.anitube.site/?s="

	for i := 1; i < len(fname); i++ {
		URL += fname[i] + "+"
	}

	fmt.Println()

	c.OnHTML(".aniItem", func(e *colly.HTMLElement) {
		href := e.ChildAttr("a", "href")
		name := e.ChildText(".aniItemNome")

		animes = append(animes, AnimeInfo{Name: name, ID: href, Index: e.Index})

		fmt.Printf("[%d] - %v.\n", e.Index+1, name)
	})

	c.Visit(URL)

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&option)

	for index, anime := range animes {
		if (index + 1) == option {
			animeSelected = anime.ID
			break
		}
	}

	utils.Clear()

	return animeSelected
}
