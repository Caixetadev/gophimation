package episode

import (
	"fmt"
	"log"

	"github.com/Caixetadev/gophimation/configs"
	"github.com/Caixetadev/gophimation/pkg/presence"
	"github.com/Caixetadev/gophimation/pkg/util"
	"github.com/gocolly/colly/v2"
)

func SelectEpisode(URL string) (string, string, string, string, string, string) {
	c := configs.Colly()

	var episodes []util.AnimeInfo
	var option int
	var episodeSelected string
	var imageAnime string
	var nameEpisode string
	var nameAnime string

	c.OnHTML(".div_video_list a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		name := e.Text

		episodes = append(episodes, util.AnimeInfo{Name: name, ID: href, Index: e.Index})

		fmt.Printf("[%d] -  %v\n", e.Index+1, name)
	})

	c.OnHTML(".divMainNomeAnime", func(h *colly.HTMLElement) {
		imageAnime = h.ChildAttr(".sub_animepage_img img", "data-src")
		nameAnime = h.ChildText(".div_anime_names h1")
	})

	if err := c.Visit(URL); err != nil {
		util.UnCache(URL)
		log.Fatalln(err)
	}

	fmt.Println("\ncoloque um numero para assistir")

	fmt.Scanln(&option)

	util.OptionIsValid(episodes, option)

	for i, ai := range episodes {
		if (i + 1) == option {
			episodeSelected = ai.ID
			nameEpisode = ai.Name
		}
	}

	var watching string

	if option < 10 {
		watching = fmt.Sprintf("Episódio %02d", option)
	} else {
		watching = fmt.Sprintf("Episódio %d", option)
	}

	presence.Presence(nameEpisode, imageAnime, "Assistindo "+nameAnime, watching, "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png")

	return episodeSelected, nameEpisode, imageAnime, "Assistindo " + nameAnime, watching, "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png"
}
