package jikan

// Anime via http://jikan.me/api/v1.2/anime/1/characters_staff
type Anime struct {
	LinkCanonical string     `json:"link-canonical"`
	Synopsis      string     `json:"synopsis"`
	Title         string     `json:"title"`
	Image         string     `json:"image"`
	TitleEnglish  string     `json:"title-english"`
	Japanese      string     `json:"japanese"`
	Type          string     `json:"type"`
	Episodes      int        `json:"episodes"`
	Status        string     `json:"status"`
	Aired         string     `json:"aired"`
	Premiered     string     `json:"premiered"`
	Broadcast     string     `json:"broadcast"`
	Producer      [][]string `json:"producer"`
	Licensor      [][]string `json:"licensor"`
	Studio        [][]string `json:"studio"`
	Source        string     `json:"source"`
	Genre         [][]string `json:"genre"`
	Duration      string     `json:"duration"`
	Rating        string     `json:"rating"`
	Score         []int      `json:"score"`
	Ranked        int        `json:"ranked"`
	Popularity    int        `json:"popularity"`
	Members       int        `json:"members"`
	Favorites     int        `json:"favorites"`
	Background    string     `json:"background"`
	Related       struct {
		Adaptation [][]string `json:"Adaptation"`
		SideStory  [][]string `json:"Side story"`
		Summary    [][]string `json:"Summary"`
	} `json:"related"`
	OpeningTheme []string `json:"opening-theme"`
	EndingTheme  []string `json:"ending-theme"`
	Character    []struct {
		Image      string `json:"image"`
		URL        string `json:"url"`
		Name       string `json:"name"`
		Role       string `json:"role"`
		VoiceActor []struct {
			Name  string `json:"name"`
			URL   string `json:"url"`
			Role  string `json:"role"`
			Image string `json:"image"`
		} `json:"voice-actor"`
	} `json:"character"`
	Staff []struct {
		Image string `json:"image"`
		Name  string `json:"name"`
		URL   string `json:"url"`
		Role  string `json:"role"`
	} `json:"staff"`
}
