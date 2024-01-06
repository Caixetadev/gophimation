package random

import (
	"github.com/Caixetadev/gophimation/pkg/episode"
	"github.com/Caixetadev/gophimation/pkg/selectVideo"
)

func Random() {
	s, _ := episode.SelectEpisode("random")
	selectVideo.SelectVideo(s)
}
