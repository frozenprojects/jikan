package jikan

import (
	"github.com/aerogo/http/client"
)

// Character via https://jikan.me/api/v1.1/character/1
type Character struct {
	Image        string `json:"image"`
	Animeography []struct {
		Name  string `json:"name"`
		Link  string `json:"link"`
		Image string `json:"image"`
	} `json:"characterography"` // TODO: Fix this to animeography
	Mangaography []struct {
		Name  string `json:"name"`
		Link  string `json:"link"`
		Image string `json:"image"`
	} `json:"mangaography"`
	MemberFavorites int    `json:"member-favorites"`
	Name            string `json:"name"`
	NameJapanese    string `json:"name-japanese"`
	About           string `json:"about"`
	VoiceActors     []struct {
		Name     string `json:"name"`
		Link     string `json:"link"`
		Image    string `json:"image"`
		Language string `json:"language"`
	} `json:"voice-actors"`
}

// GetCharacter returns a character object for the given ID.
func GetCharacter(id string) (*Character, error) {
	character := &Character{}
	_, err := client.Get(CharacterEndpoint + "/character/" + id).EndStruct(character)
	return character, err
}
