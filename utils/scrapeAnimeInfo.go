package utils

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type AnimeInfo struct {
	Index int
	Name  string
	ID    string
}

var animes []AnimeInfo

func ScrapeAnimeInfo(e *colly.HTMLElement) []AnimeInfo {
	href := e.ChildAttr("a", "href")
	name := e.ChildText(".aniItemNome")

	animes = append(animes, AnimeInfo{Name: name, ID: href, Index: e.Index})

	fmt.Printf("[%d] - %v.\n", e.Index+1, name)

	return animes
}
