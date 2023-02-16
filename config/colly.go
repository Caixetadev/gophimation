package config

import (
	"github.com/gocolly/colly/v2"
)

func Colly() *colly.Collector {
	c := colly.NewCollector(
		colly.CacheDir("./cache"),
		colly.AllowedDomains("animefire.net"),
	)

	return c
}
