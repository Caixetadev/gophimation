package search

import (
	"fmt"
	"log"
	"os"

	"github.com/Caixetadev/gophimation/configs"
	mostwatched "github.com/Caixetadev/gophimation/pkg/mostWatched"
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

	fname := os.Args
	var option int
	var animes []util.AnimeInfo
	var animeSelected string

	var hasArgs = len(fname) == 1

	if hasArgs {
		URL := mostwatched.MostWatched()

		return URL
	}

	URL := "https://animefire.net/pesquisar/"

	for i := 1; i < len(fname); i++ {
		URL += fname[i] + "-"
	}

	c.OnHTML(".card", func(e *colly.HTMLElement) {
		animes = util.ScrapeAnimeInfo(e)
	})

	if err := c.Visit(URL); err != nil {
		log.Fatalln(err)
	}

	if len(animes) == 0 {
		util.Clear()
		log.Fatal("Não foi possivel achar o anime")
	}

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&option)

	util.OptionIsValid(animes, option)

	for index, anime := range animes {
		if (index + 1) == option {
			animeSelected = anime.ID
			break
		}
	}

	util.Clear()

	return animeSelected
}
