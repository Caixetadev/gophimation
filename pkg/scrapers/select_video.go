package scrapers

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/Caixetadev/gophimation/config"
	"github.com/Caixetadev/gophimation/internal/utils"
	"github.com/Caixetadev/gophimation/pkg/constants"
	"github.com/gocolly/colly"
	"github.com/peterbourgon/diskv/v3"
)

type PlayerInfo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// SelectVideo does the search for the url of the video
func SelectVideo(ep string) *PlayerInfo {
	fmt.Println(ep)
	var urlPlayer []PlayerInfo

	// Initialize a new diskv store, rooted at "cache-dir", with a 10MB cache.
	d := diskv.New(diskv.Options{
		BasePath:     utils.GetHomeDir("gophimation"),
		Transform:    func(s string) []string { return []string{} },
		CacheSizeMax: 10 * 1024 * 1024,
	})

	key := strings.ReplaceAll(strings.ReplaceAll(ep, "-", "_"), "/", "_")

	data, _ := d.Read(key)

	if len(data) > 1 {
		err := json.Unmarshal(data, &urlPlayer)

		if err != nil {
			log.Fatal(err)
		}

		return &PlayerInfo{Name: urlPlayer[0].Name, Url: urlPlayer[0].Url}
	}

	c := config.Colly()

	iframeURL, nameAnimeAndEpisode := utils.GetIframe(constants.URL_BASE + ep)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", constants.URL_BASE)
	})

	c.OnHTML("script:nth-of-type(4)", func(h *colly.HTMLElement) {
		re := regexp.MustCompile(`"file":"([^"]+)"`)
		match := re.FindStringSubmatch(h.Text)

		if len(match) > 1 {
			urlPlayer = append(urlPlayer, PlayerInfo{Name: nameAnimeAndEpisode, Url: strings.ReplaceAll(match[1], "\\", "")})
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(err)
	})

	if err := c.Visit(iframeURL); err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := json.Marshal(urlPlayer)

	if err != nil {
		log.Fatal(err)
	}

	// Write the data to the cache.
	if err := d.Write(key, jsonBytes); err != nil {
		log.Fatal(err)
	}

	return &PlayerInfo{Name: urlPlayer[0].Name, Url: urlPlayer[0].Url}
}
