package search

import (
	"fmt"
	"log"
	"os"

	"github.com/Caixetadev/gophimation/configs"
	mostWatched "github.com/Caixetadev/gophimation/pkg/mostWatched"
	"github.com/Caixetadev/gophimation/pkg/util"
	"github.com/gocolly/colly/v2"
)

type AnimeInfo struct {
	Index int
	Name  string
	ID    string
}

// Search does the search for the anime
func Search() string {
	c := configs.Colly()

	flags := os.Args
	var option int
	var anime []util.AnimeInfo
	var animeSelected string

	var hasArgs = len(flags) == 1

	if hasArgs {
		URL := mostWatched.MostWatched()

		return URL
	}

	URL := "https://animefire.net/pesquisar/"

	for i := 1; i < len(flags); i++ {
		URL += flags[i] + "-"
	}

	c.OnHTML(".card", func(e *colly.HTMLElement) {
		anime = util.ScrapeAnimeInfo(e)
	})

	if err := c.Visit(URL); err != nil {
		log.Fatalln(err)
	}

	if len(anime) == 0 {
		util.Clear()
		log.Fatal("Não foi possivel achar o anime")
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
