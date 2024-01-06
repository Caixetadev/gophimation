package scrapers

import (
	"github.com/Caixetadev/gophimation/internal/utils"
)

func Random() {
	episodeSelected, nextEpisode := SelectEpisode("random")

	videoSelected := SelectVideo(episodeSelected)
	if nextEpisode != nil {
		go SelectVideo(*nextEpisode)
	}

	utils.PlayVideo(videoSelected.Url, videoSelected.Name)
}
