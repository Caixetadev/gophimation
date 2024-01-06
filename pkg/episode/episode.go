package episode

import (
	"fmt"
	"log"
	"strings"

	"github.com/Caixetadev/gophimation/pkg/configs"
	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/Caixetadev/gophimation/pkg/models"
	"github.com/Caixetadev/gophimation/pkg/presence"
	"github.com/Caixetadev/gophimation/pkg/util"
	"github.com/gocolly/colly"
)

func SelectEpisode(URL string) (string, string) {
	var selectedOption int
	var episodes []models.Anime

	var nameAnime string
	var image string

	c := configs.Colly()

	c.OnHTML(".infos_left .anime-info", func(h *colly.HTMLElement) {
		nameAnime = h.ChildText("h2")

		if URL == "random" {
			fmt.Printf("O anime random é: %s\n\n", nameAnime)
		}
	})

	c.OnHTML("#episodesList .list-group-item-action", func(h *colly.HTMLElement) {
		fmt.Printf("[%02d] - %v\n", h.Index+1, h.ChildText("a h3"))

		episodes = append(episodes, models.Anime{Name: h.ChildText("a h3"), URL: strings.TrimPrefix(h.ChildAttr("a", "href"), constants.URL_BASE)})
	})

	c.OnHTML("main.container", func(h *colly.HTMLElement) {
		image = h.ChildAttr(".infos-img img", "src")
	})

	if err := c.Visit(constants.URL_BASE + URL); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\ncoloque um numero para assistir to aqui")

	fmt.Scanln(&selectedOption)

	util.OptionIsValid(episodes, selectedOption)

	util.Clear()

	fmt.Println("Carregando...")

	go presence.Presence("https:"+image, nameAnime, fmt.Sprintf("Episódio %02d", selectedOption), "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png")

	return episodes[selectedOption-1].URL, episodes[(selectedOption+1)-1].URL
}
