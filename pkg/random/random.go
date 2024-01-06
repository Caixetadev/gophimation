package random

import (
	"github.com/Caixetadev/gophimation/pkg/episode"
	"github.com/Caixetadev/gophimation/pkg/selectVideo"
	"github.com/Caixetadev/gophimation/pkg/util"
)

func Random() {
	episodeSelected, nextEpisode := episode.SelectEpisode("random")

	videoSelected := selectVideo.SelectVideo(episodeSelected)
	if nextEpisode != nil {
		go selectVideo.SelectVideo(*nextEpisode)
	}

	util.PlayVideo(videoSelected.Url, videoSelected.Name)
}
