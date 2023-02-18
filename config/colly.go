package config

import (
	"github.com/gocolly/colly/v2"
)

func Colly() *colly.Collector {
	c := colly.NewCollector(
		colly.CacheDir("./cache"),
		colly.AllowedDomains("animefire.net", "www.blogger.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"),
	)

	return c
}
