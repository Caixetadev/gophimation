package selectVideo

import (
	"encoding/json"
	"io"
	"log"
	"regexp"
	"strings"

	"github.com/Caixetadev/gophimation/config"
	"github.com/Caixetadev/gophimation/utils"
	"github.com/gocolly/colly/v2"
)

type VideoSource struct {
	Src   string `json:"src"`
	Label string `json:"label"`
}

type ApiResponse struct {
	Data []VideoSource `json:"data"`
}

func SelectVideo(ep, nameAnime string) {
	c := config.Colly()

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

		utils.PlayVideo(response.Data[len(response.Data)-1].Src, nameAnime)
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

				utils.PlayVideo(strings.Replace(urlstring, `"`, "", -1), nameAnime)
			})
		}
	})

	if err := c.Visit(ep); err != nil {
		log.Fatalln(err)
	}
}
