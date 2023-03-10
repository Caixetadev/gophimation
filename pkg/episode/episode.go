package episode

import (
	"fmt"
	"strings"

	"github.com/Caixetadev/gophimation/pkg/configs"
	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/Caixetadev/gophimation/pkg/models"
	"github.com/Caixetadev/gophimation/pkg/presence"
	"github.com/Caixetadev/gophimation/pkg/util"
	"github.com/gocolly/colly"
)

func SelectEpisode(URL string) string {
	var selectedOption int
	var episodes []models.Anime

	var nameAnime2 string

	var image2 string

	c := configs.Colly()

	c.OnHTML("#episodesList .list-group-item-action", func(h *colly.HTMLElement) {
		episode := h.ChildText("a h3")
		urlAnime := h.ChildAttr("a", "href")

		episodes = append(episodes, models.Anime{Name: episode, URL: strings.TrimPrefix(urlAnime, constants.URL_BASE)})
	})

	c.OnHTML("main.container", func(h *colly.HTMLElement) {
		image := h.ChildAttr(".infos-img img", "src")

		image2 = image
	})

	c.OnHTML(".infos_left .anime-info", func(h *colly.HTMLElement) {
		nameAnime := h.ChildText("h2")

		nameAnime2 = nameAnime
	})

	c.Visit(constants.URL_BASE + URL)

	animeResponse := models.AnimeResponse{
		Anime:    models.Anime{Name: nameAnime2, URL: image2},
		Episodes: episodes,
	}

	util.Clear()

	if URL == "random" {
		fmt.Printf("O anime random é %s\n", animeResponse.Anime.Name)
	}

	fmt.Println()

	for i, item := range animeResponse.Episodes {
		fmt.Printf("[%02d] - %v\n", i+1, item.Name)
	}

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&selectedOption)

	util.OptionIsValid(animeResponse.Episodes, selectedOption)

	var watching string

	if selectedOption < 10 {
		watching = fmt.Sprintf("Episódio %02d", selectedOption)
	} else {
		watching = fmt.Sprintf("Episódio %d", selectedOption)
	}

	util.Clear()

	fmt.Println("Carregando...")

	presence.Presence("Caixeta", "https:"+animeResponse.Anime.URL, animeResponse.Anime.Name, watching, "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png")

	return animeResponse.Episodes[selectedOption-1].URL
}
