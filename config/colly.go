package config

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/gocolly/colly"
)

func CollyPastebin() string {
	c := colly.NewCollector(
		colly.AllowedDomains("pastebin.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"),
	)

	var cookie string
	c.OnHTML(".content .post-view .de1", func(e *colly.HTMLElement) {
		cookie = e.Text
	})

	c.Visit("https://pastebin.com/9iNGXsDt")

	return cookie
}

func Colly() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("betteranime.net"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"),
	)

	cookieJar, err := cookiejar.New(nil)

	if err != nil {
		log.Fatal(err)
	}

	c.SetCookieJar(cookieJar)

	url, err := url.Parse("https://betteranime.net")

	if err != nil {
		log.Fatal(err)
	}

	cookiePastebin := CollyPastebin()

	cookie := &http.Cookie{
		Name:   "betteranime_session",
		Value:  cookiePastebin,
		Domain: "betteranime.net",
		Path:   "/",
	}

	cookie2 := &http.Cookie{
		Name:   "BetterQuality",
		Value:  "eyJpdiI6ImVqbEcwT0dQZWNNNjFuK0NwZUVjMnc9PSIsInZhbHVlIjoib3RQZFF0TEZGcTZwM2pjRFJ1aU8yOWRLOW5ORFh4M1pkSzdEblZ0T2IrMmxTSGgwaHNCUHVrQTZ1MDBEbkRkZy93aHIyak9xVWh1Wmc5K05BRUNqYUMrZzIvNzY4elpwNDRUMWplN2ZOMXNnd3k0QWgwb3p3SFZYYWF5S0g3RjAiLCJtYWMiOiJiYWViNTA2NTA4NTY1NTRiZmY0Yjg1Y2U2MzI4ZTdlZGYxNDUzNmU3NGMwZGRmMTM5MTQ4OTJmMGNjODQ2MjIwIiwidGFnIjoiIn0%3D",
		Domain: "betteranime.net",
		Path:   "/",
	}

	cookieJar.SetCookies(url, []*http.Cookie{cookie, cookie2})

	return c
}
