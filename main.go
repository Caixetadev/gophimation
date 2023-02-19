package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/Caixetadev/gophimation/config"
	"github.com/Caixetadev/gophimation/constants"
	"github.com/Caixetadev/gophimation/episode"
	"github.com/Caixetadev/gophimation/presence"
	"github.com/Caixetadev/gophimation/utils"
	"github.com/gocolly/colly/v2"
)

func init() {
	presence.Presence("Caixeta", "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png", "Explorando Animes", "Encontre seu próximo anime favorito <3", "")
	utils.Clear()

	_, error := os.Stat(constants.FILE_NAME)

	if os.IsNotExist(error) {
		utils.CreateFile(constants.FILE_NAME)
	} else {
		utils.ReadFile(constants.FILE_NAME)
	}
}

type VideoSource struct {
	Src   string `json:"src"`
	Label string `json:"label"`
}

type ApiResponse struct {
	Data []VideoSource `json:"data"`
}

func main() {
	c := config.Colly()

	ep := episode.SelectEpisode()

	client := config.Http()

	var triggerHTML bool

	c.OnHTML("#my-video", func(e *colly.HTMLElement) {
		URL := e.Attr("data-video-src")

		res, err := client.Get(URL)

		if err != nil {
			log.Fatal("error")
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()

		var response ApiResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Fatalln(err)
		}

		utils.Clear()

		utils.PlayVideo(response.Data[len(response.Data)-1].Src)
	})

	c.OnHTML("#div_video iframe", func(h *colly.HTMLElement) {
		URL := h.Attr("src")

		triggerHTML = true

		if err := c.Visit(URL); err != nil {
			log.Fatalln(err)
		}
	})

	c.OnResponse(func(r *colly.Response) {
		if triggerHTML {
			c.OnHTML("html", func(e *colly.HTMLElement) {
				res := regexp.MustCompile(`"https://rr[\S]+?"`)
				url := res.FindAllStringSubmatch(e.Text, -1)

				urlstring := strings.Join(res.FindAllStringSubmatch(e.Text, -1)[len(url)-1], "")

				utils.Clear()

				utils.PlayVideo(strings.Replace(urlstring, `"`, "", -1))
			})
		}
	})

	if err := c.Visit(ep); err != nil {
		log.Fatalln(err)
	}
}
