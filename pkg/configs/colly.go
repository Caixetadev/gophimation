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
		Name:   "betteranime_session",
		Value:  "eyJpdiI6IjhWUjR3aGp6YTZNK2ZOcjdFRGVwdlE9PSIsInZhbHVlIjoicUM2K1JDVFQvQTVBZm90bmxPZDRMMHdEZmdRdGQvbFNROXBKYmdNR3NFTU1yNTNUTElNdDlQUlArRmtqS2JtNVJqUndtWFE4TlYwU0Q1QWxjNjNHQ1RNWTNyRjZHbjIyT0VWanBpaTlsRDBCU2RXcitzSm92SEdOMnJqVnlucFciLCJtYWMiOiJmNTJlNDZkNjBkNDk1ZDkxZTA0OTg1YThkNDMzODQ5M2ZkZmRlZWVkNDMyODY5OWYzYTdjMTFjMzFlYTE0MTJkIiwidGFnIjoiIn0%3D",
		Domain: "betteranime.net",
		Path:   "/",
	}

	cookie2 := &http.Cookie{
		Name:   "BetterQuality",
		Value:  "eyJpdiI6ImtRS3lmckhma3FzVFZiV1JDTmpINHc9PSIsInZhbHVlIjoicU9WWTVSdmR0V3JrSnJRbjRSK1FFVnIxNXo4dVNSYjltaldJWmpzK1krSm56eWJyRWVBRXk3OEVHVng2Z1docSIsIm1hYyI6IjE5MGQ3OGQwOGYzZmMzNWMyMTg5M2Q2ZWM4MjNmYTZlOTU2NjFkNzVkMTAzOWI0NWE5MWU5ZjY1YTM4YmRlOTciLCJ0YWciOiIifQ==",
		Domain: "betteranime.net",
		Path:   "/",
	}

	cookieJar.SetCookies(url, []*http.Cookie{cookie, cookie2})

	return c
}
