package random

import (
	"github.com/Caixetadev/gophimation/pkg/episode"
	"github.com/Caixetadev/gophimation/pkg/selectVideo"
)

func Random() {
	selectVideo.SelectVideo(episode.SelectEpisode("random"))
}
