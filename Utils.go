package jikan

import "regexp"

var characterIDRegEx = regexp.MustCompile(`myanimelist.net/character/(\d+)/.*`)

// GetCharacterIDFromURL retrieves the character ID from the character URL.
func GetCharacterIDFromURL(url string) string {
	matches := characterIDRegEx.FindStringSubmatch(url)

	if len(matches) > 1 {
		return matches[1]
	}

	return ""
}
