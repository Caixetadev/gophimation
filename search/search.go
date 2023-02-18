package search

import (
	"fmt"
	"log"
	"os"

	"github.com/Caixetadev/gophimation/config"
	mostwatched "github.com/Caixetadev/gophimation/mostWatched"
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
	var animes []utils.AnimeInfo
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
		animes = utils.ScrapeAnimeInfo(e)
	})

	if err := c.Visit(URL); err != nil {
		log.Fatalln(err)
	}

	if len(animes) == 0 {
		utils.Clear()
		log.Fatal("Não foi possivel achar o anime")
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
