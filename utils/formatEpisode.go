package utils

import "fmt"

func FormatEpisode(episode int) string {
	var watching string

	if episode < 10 {
		watching = fmt.Sprintf("Episódio %02d", episode)
	} else {
		watching = fmt.Sprintf("Episódio %d", episode)
	}

	return watching
}
