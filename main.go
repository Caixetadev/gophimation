package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/Caixetadev/gophimation/config"
	"github.com/Caixetadev/gophimation/episode"
	"github.com/Caixetadev/gophimation/utils"
	"github.com/gocolly/colly/v2"
)

const FILE_NAME = "dataUser.json"

type UserConfig struct {
	Name string `json:"name"`
}

func init() {
	var name string

	utils.Clear()

	_, error := os.Stat(FILE_NAME)

	if os.IsNotExist(error) {
		fmt.Println("Qual o seu nome?")

		fmt.Scanln(&name)

		file, err := os.Create(FILE_NAME)

		data := UserConfig{
			Name: name,
		}

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		user, _ := json.Marshal(data)

		_ = os.WriteFile(FILE_NAME, user, 0644)

		if err != nil {
			log.Fatal(err)
		}

		utils.Clear()
	} else {
		data, err := os.ReadFile(FILE_NAME)

		var user UserConfig

		if errUnmarshal := json.Unmarshal(data, &user); errUnmarshal != nil {
			log.Fatalln(errUnmarshal)
		}

		utils.Greeting(user.Name)

		fmt.Println()

		if err != nil {
			log.Panicf("failed reading data from file: %s", err)
		}
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
