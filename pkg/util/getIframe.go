package util

import (
	"fmt"
	"log"

	"github.com/Caixetadev/gophimation/pkg/configs"
	"github.com/gocolly/colly"
)

func GetIframe(URL string) (string, string) {
	c := configs.Colly()

	var iframe string
	var nameAnimeAndEpisode string

	c.OnHTML(".anime-title", func(h *colly.HTMLElement) {
		nameAnime := h.ChildText("h2 a")
		nameEpisode := h.ChildText("h3")

		nameAnimeAndEpisode = fmt.Sprintf("%s - %s", nameAnime, nameEpisode)
	})

	c.OnHTML("iframe", func(h *colly.HTMLElement) {
		iframe = h.Attr("src")
	})

	if err := c.Visit(URL); err != nil {
		log.Fatal(err)
	}

	return iframe, nameAnimeAndEpisode
}
