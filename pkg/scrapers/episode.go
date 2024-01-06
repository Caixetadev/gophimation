package scrapers

import (
	"fmt"
	"log"
	"strings"

	"github.com/Caixetadev/gophimation/config"
	"github.com/Caixetadev/gophimation/internal/entity"
	"github.com/Caixetadev/gophimation/internal/presence"
	"github.com/Caixetadev/gophimation/internal/utils"
	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/gocolly/colly"
)

func SelectEpisode(URL string) (string, *string) {
	var selectedOption int
	var episodes []entity.Anime

	var nameAnime string
	var image string

	c := config.Colly()

	c.OnHTML(".infos_left .anime-info", func(h *colly.HTMLElement) {
		nameAnime = h.ChildText("h2")

		if URL == "random" {
			fmt.Printf("O anime random é: %s\n\n", nameAnime)
		}
	})

	c.OnHTML("#episodesList .list-group-item-action", func(h *colly.HTMLElement) {
		fmt.Printf("[%02d] - %v\n", h.Index+1, h.ChildText("a h3"))

		episodes = append(episodes, entity.Anime{Name: h.ChildText("a h3"), URL: strings.TrimPrefix(h.ChildAttr("a", "href"), constants.URL_BASE)})
	})

	c.OnHTML("main.container", func(h *colly.HTMLElement) {
		image = h.ChildAttr(".infos-img img", "src")
	})

	if err := c.Visit(constants.URL_BASE + URL); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\ncoloque um numero para assistir to aqui")

	fmt.Scanln(&selectedOption)

	utils.OptionIsValid(episodes, selectedOption)

	utils.Clear()

	fmt.Println("Carregando...")

	go presence.Presence("https:"+image, nameAnime, fmt.Sprintf("Episódio %02d", selectedOption), "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png")

	if selectedOption == len(episodes) {
		return episodes[selectedOption-1].URL, nil
	}

	return episodes[selectedOption-1].URL, &episodes[(selectedOption+1)-1].URL
}
