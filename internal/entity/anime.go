package entity

type Anime struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type AnimeResponse struct {
	Anime    Anime   `json:"anime"`
	Episodes []Anime `json:"episodes"`
}
