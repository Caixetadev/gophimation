package episode

import (
	"fmt"
	"log"

	"github.com/Caixetadev/gophimation/config"
	"github.com/Caixetadev/gophimation/presence"
	"github.com/Caixetadev/gophimation/search"
	"github.com/Caixetadev/gophimation/utils"
	"github.com/gocolly/colly/v2"
)

func SelectEpisode() string {
	c := config.Colly()

	var episodes []utils.AnimeInfo
	var option int
	var episodeSelected string
	var imageAnime string
	var nameEpisode string
	var nameAnime string

	c.OnHTML(".div_video_list a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		name := e.Text

		episodes = append(episodes, utils.AnimeInfo{Name: name, ID: href, Index: e.Index})

		fmt.Printf("[%d] -  %v\n", e.Index+1, name)
	})

	c.OnHTML(".divMainNomeAnime", func(h *colly.HTMLElement) {
		imageAnime = h.ChildAttr(".sub_animepage_img img", "data-src")
		nameAnime = h.ChildText(".div_anime_names h1")
	})

	URL := search.Search()

	if err := c.Visit(URL); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&option)

	utils.OptionIsValid(episodes, option)

	for i, ai := range episodes {
		if (i + 1) == option {
			episodeSelected = ai.ID
			nameEpisode = ai.Name
		}
	}

	watching := fmt.Sprintf("Episode %s", nameAnime)

	presence.Presence(nameEpisode, imageAnime, nameAnime, watching)

	return episodeSelected
}
