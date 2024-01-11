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

const userPromptEpisode = "\nColoque um numero para assistir"

// SelectEpisode performs web scraping on the episodes page, allowing the user to choose the desired episode.
func SelectEpisode(URL string) (*string, int, []entity.Anime) {
	c := config.Colly()

	var selectedOption int
	var episodes []entity.Anime
	var nameAnime string
	var image string

	setCollyCallbacksEpisodes(c, &nameAnime, &episodes, &image, URL)

	if err := c.Visit(constants.URL_BASE + URL); err != nil {
		log.Fatal(err)
	}

	selectedOption = utils.GetUserInput(userPromptEpisode)

	utils.OptionIsValid(episodes, selectedOption)

	utils.Clear()

	fmt.Println("Carregando...")

	updatePresence(image, nameAnime, selectedOption)

	if selectedOption == len(episodes) {
		return nil, selectedOption - 1, episodes
	}

	return &episodes[(selectedOption+1)-1].URL, selectedOption - 1, episodes
}

func setCollyCallbacksEpisodes(c *colly.Collector, nameAnime *string, episodes *[]entity.Anime, image *string, URL string) {
	c.OnHTML(".infos_left .anime-info", func(h *colly.HTMLElement) {
		*nameAnime = h.ChildText("h2")

		if URL == "random" {
			fmt.Printf("O anime random é: %s\n\n", *nameAnime)
		}
	})

	c.OnHTML("#episodesList .list-group-item-action", func(h *colly.HTMLElement) {
		fmt.Printf("[%02d] - %v\n", h.Index+1, h.ChildText("a h3"))

		*episodes = append(*episodes, entity.Anime{Name: h.ChildText("a h3"), URL: strings.TrimPrefix(h.ChildAttr("a", "href"), constants.URL_BASE)})
	})

	c.OnHTML("main.container", func(h *colly.HTMLElement) {
		*image = h.ChildAttr(".infos-img img", "src")
	})
}

func updatePresence(image string, nameAnime string, selectedOption int) {
	go presence.Presence("https:"+image, nameAnime, fmt.Sprintf("Episódio %02d", selectedOption), "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png")
}
