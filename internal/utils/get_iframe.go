package utils

import (
	"fmt"
	"log"

	"github.com/Caixetadev/gophimation/config"
	"github.com/gocolly/colly"
)

func GetIframe(URL string) (string, string) {
	c := config.Colly()

	var iframe string
	var nameAnimeAndEpisode string

	c.OnHTML(".anime-title", func(h *colly.HTMLElement) {
		nameAnimeAndEpisode = fmt.Sprintf("%s - %s", h.ChildText("h2 a"), h.ChildText("h3"))
	})

	c.OnHTML("iframe", func(h *colly.HTMLElement) {
		iframe = h.Attr("src")
	})

	if err := c.Visit(URL); err != nil {
		log.Fatal(err)
	}

	return iframe, nameAnimeAndEpisode
}
