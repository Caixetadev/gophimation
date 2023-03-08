package models

type AnimeInfo struct {
	Index int
	Name  string `json:"name"`
	ID    string `json:"url"`
}
