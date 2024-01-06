package scrapers

import (
	"fmt"
	"log"
	"strings"

	"github.com/Caixetadev/gophimation/config"
	"github.com/Caixetadev/gophimation/internal/entity"
	"github.com/Caixetadev/gophimation/internal/utils"
	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/gocolly/colly"
)

const userPromptMostWatched = "\nColoque um numero para assistir"

// MostWatched prints the most viewed anime of the week and returns its url
func MostWatched() string {
	var selectedOption int

	c := config.Colly()

	var anime []entity.Anime

	setCollyCallbacksMostWatched(c, &anime)

	if err := c.Visit(constants.URL_BASE); err != nil {
		log.Fatal(err)
	}

	selectedOption = utils.GetUserInput(userPromptMostWatched)

	utils.OptionIsValid(anime, selectedOption)

	utils.Clear()

	return anime[selectedOption-1].URL
}

func setCollyCallbacksMostWatched(c *colly.Collector, anime *[]entity.Anime) {
	c.OnHTML(".highlights .highlight-card .highlight-body", func(h *colly.HTMLElement) {
		fmt.Printf("[%02d] - %v\n", h.Index+1, h.ChildText(".highlight-title h3"))

		*anime = append(*anime, entity.Anime{URL: strings.TrimPrefix(h.ChildAttr("a", "href"), constants.URL_BASE)})
	})
}
