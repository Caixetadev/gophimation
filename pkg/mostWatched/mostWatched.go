package mostWatched

import (
	"fmt"
	"log"
	"strings"

	"github.com/Caixetadev/gophimation/pkg/configs"
	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/Caixetadev/gophimation/pkg/models"
	"github.com/Caixetadev/gophimation/pkg/util"
	"github.com/gocolly/colly"
)

// MostWatched prints the most viewed anime of the week and returns its url
func MostWatched() string {
	var selectedOption int

	c := configs.Colly()

	var anime []models.Anime

	c.OnHTML(".highlights .highlight-card .highlight-body", func(h *colly.HTMLElement) {
		fmt.Printf("[%02d] - %v\n", h.Index+1, h.ChildText(".highlight-title h3"))

		anime = append(anime, models.Anime{URL: strings.TrimPrefix(h.ChildAttr("a", "href"), constants.URL_BASE)})
	})

	if err := c.Visit(constants.URL_BASE); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&selectedOption)

	util.OptionIsValid(anime, selectedOption)

	util.Clear()

	return anime[selectedOption-1].URL
}
