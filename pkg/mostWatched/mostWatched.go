package mostwatched

import (
	"fmt"
	"strings"

	"github.com/Caixetadev/gophimation/pkg/configs"
	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/Caixetadev/gophimation/pkg/models"
	"github.com/Caixetadev/gophimation/pkg/util"
	"github.com/gocolly/colly"
)

// MostWatched prints the most viewed anime of the week and returns its url
func MostWatched() string {
	var option int

	c := configs.Colly()

	var anime []models.Anime

	c.OnHTML(".highlights .highlight-card .highlight-body", func(h *colly.HTMLElement) {
		urlAnime := h.ChildAttr("a", "href")
		name := h.ChildText(".highlight-title h3")

		anime = append(anime, models.Anime{Name: name, URL: strings.TrimPrefix(urlAnime, constants.URL_BASE)})
	})

	c.Visit(constants.URL_BASE)

	for i, item := range anime {
		fmt.Printf("[%02d] - %v\n", i+1, item.Name)
	}

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&option)

	util.OptionIsValid(anime, option)

	util.Clear()

	return anime[option-1].URL
}
