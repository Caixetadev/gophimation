package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

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

		json.Unmarshal(data, &user)

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
	Data     []VideoSource `json:"data"`
	Response struct {
		Status string `json:"status"`
		Text   string `json:"text"`
	} `json:"response"`
}

func main() {
	c := config.Colly()

	ep := episode.SelectEpisode()

	c.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	c.OnHTML("#my-video", func(e *colly.HTMLElement) {
		URL := e.Attr("data-video-src")

		res, err := http.Get(URL)

		if err != nil {
			log.Fatal("error")
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()

		var response ApiResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Fatalln(err)
		}

		cmd := exec.Command("mpv", response.Data[len(response.Data)-1].Src, "--demuxer-max-bytes=1G", "--no-terminal", "--fs", "video")
		cmd.Stdout = os.Stdout
		cmd.Run()
	})

	c.Visit(ep)
}
