package configs

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/gocolly/colly"
)

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

	cookie := &http.Cookie{
		Name:   "BetterQuality",
		Value:  "eyJpdiI6IlhNbnFZbVMwV2pHK2xNU1FBTUd1dEE9PSIsInZhbHVlIjoibmh4bVVBT1h2TVBJaEFXSzI3Q2xlUVlNMWxZTlpOSEZtUHBtcVlxa0hDSlRteHpQS083ZmpDQ2xramFOc29NciIsIm1hYyI6Ijg2NzA3M2ZkZDQ4MWY3YzgxNjIzZDQxZjViYzJhNGRjNGY5MTk3NzVhNzI2NWJkYzUxOGNiMDdmYzQyNTA2NGUiLCJ0YWciOiIifQ%3D%3D",
		Domain: "betteranime.net",
		Path:   "/",
	}

	cookieJar.SetCookies(url, []*http.Cookie{cookie})

	return c
}
