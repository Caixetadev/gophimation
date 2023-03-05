package util

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

// AnimeInfo is a struct that stores information about an anime, such as its index, name and ID.
type AnimeInfo struct {
	Index int
	Name  string
	ID    string
}

// Slice da struct AnimeInfo
var animes []AnimeInfo

// ScrapeAnimeInfo is a function that takes an HTML element and extracts the relevant information, such as the name of the anime, its URL and index.
// And returns this information in slice format from struct AnimeInfo
func ScrapeAnimeInfo(e *colly.HTMLElement) []AnimeInfo {
	href := e.ChildAttr("a", "href")
	name := e.ChildText(".text-block h3")

	// Creates an AnimeInfo object with the extracted information and adds it to the "animes" slice.
	animes = append(animes, AnimeInfo{Name: name, ID: href, Index: e.Index})

	// Prints the index and name of the anime to standard output.
	fmt.Printf("[%d] - %v.\n", e.Index+1, name)

	// Returns the slice "animes".
	return animes
}
