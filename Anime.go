package jikan

import (
	"encoding/json"
	"strings"

	"github.com/aerogo/http/client"
)

// Anime via http://jikan.me/api/v2.0-dev/anime/1/characters_staff
type Anime struct {
	MalID         int         `json:"mal_id"`
	LinkCanonical string      `json:"link_canonical"`
	Title         string      `json:"title"`
	TitleEnglish  string      `json:"title_english"`
	TitleJapanese string      `json:"title_japanese"`
	TitleSynonyms interface{} `json:"title_synonyms"`
	ImageURL      string      `json:"image_url"`
	Type          string      `json:"type"`
	Source        string      `json:"source"`
	Episodes      int         `json:"episodes"`
	Status        string      `json:"status"`
	Aired         string      `json:"aired"`
	Duration      string      `json:"duration"`
	Rating        string      `json:"rating"`
	Score         float64     `json:"score"`
	ScoredBy      int         `json:"scored_by"`
	Rank          int         `json:"rank"`
	Popularity    int         `json:"popularity"`
	Members       int         `json:"members"`
	Favorites     int         `json:"favorites"`
	Synopsis      string      `json:"synopsis"`
	Background    string      `json:"background"`
	Related       struct {
		Adaptation []struct {
			MalID int    `json:"mal_id"`
			URL   string `json:"url"`
			Title string `json:"title"`
		} `json:"Adaptation"`
		SideStory []struct {
			MalID int    `json:"mal_id"`
			URL   string `json:"url"`
			Title string `json:"title"`
		} `json:"Side story"`
		Summary []struct {
			MalID int    `json:"mal_id"`
			URL   string `json:"url"`
			Title string `json:"title"`
		} `json:"Summary"`
	} `json:"related"`
	Premiered string `json:"premiered"`
	Broadcast string `json:"broadcast"`
	Producer  []struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"producer"`
	Licensor []struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"licensor"`
	Studio []struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"studio"`
	Genre []struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"genre"`
	OpeningTheme []string `json:"opening_theme"`
	EndingTheme  []string `json:"ending_theme"`
	Character    []struct {
		ImageURL   string `json:"image_url"`
		MalID      int    `json:"mal_id"`
		URL        string `json:"url"`
		Name       string `json:"name"`
		Role       string `json:"role"`
		VoiceActor []struct {
			MalID    int    `json:"mal_id"`
			Name     string `json:"name"`
			URL      string `json:"url"`
			Language string `json:"language"`
			ImageURL string `json:"image_url"`
		} `json:"voice_actor"`
	} `json:"character"`
	Staff []struct {
		ImageURL string `json:"image_url"`
		MalID    int    `json:"mal_id"`
		URL      string `json:"url"`
		Name     string `json:"name"`
		Role     string `json:"role"`
	} `json:"staff"`
}

// GetAnime returns an anime object for the given ID.
func GetAnime(id string) (*Anime, error) {
	anime := &Anime{}
	response, err := client.Get(Endpoint + "/anime/" + id + "/characters_staff").End()

	if err != nil {
		return nil, err
	}

	jsonString := response.String()
	jsonString = strings.Replace(jsonString, `"premiered":[]`, `"premiered":""`, 1)
	jsonString = strings.Replace(jsonString, `"broadcast":[]`, `"broadcast":""`, 1)
	jsonString = strings.Replace(jsonString, `"related":[]`, `"related":null`, 1)

	err = json.Unmarshal([]byte(jsonString), anime)
	return anime, err
}
