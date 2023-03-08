package selectVideo

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Caixetadev/gophimation/pkg/util"
)

// SelectVideo does the search for the url of the video
func SelectVideo(ep, nameAnime string) {
	resp, err := http.Get("http://localhost:8000/player/" + ep)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var anime []util.AnimeInfo
	err = json.NewDecoder(resp.Body).Decode(&anime)

	if err != nil {
		log.Fatal(err)
	}

	util.PlayVideo(anime[0].ID, anime[0].Name)
}
