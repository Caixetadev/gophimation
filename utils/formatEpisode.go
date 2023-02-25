package utils

import "fmt"

// FormatEpisode is a function that takes an integer episode number and formats it as a string.
// It checks if the episode number is less than 10 and adds a leading zero for format consistency.
// The function returns the formatted episode string.
func FormatEpisode(episode int) string {
	var watching string

	if episode < 10 {
		watching = fmt.Sprintf("Episódio %02d", episode)
	} else {
		watching = fmt.Sprintf("Episódio %d", episode)
	}

	return watching
}
