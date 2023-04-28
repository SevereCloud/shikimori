package shikimori

type Anime struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	Russian       string     `json:"russian"`
	Image         AnimeImage `json:"image"`
	URL           string     `json:"url"`
	Kind          string     `json:"kind"`
	Score         string     `json:"score"`
	Status        string     `json:"status"`
	Episodes      int        `json:"episodes"`
	EpisodesAired int        `json:"episodes_aired"`
	AiredOn       string     `json:"aired_on"`
	ReleasedOn    *string    `json:"released_on"`
}

type AnimeImage struct {
	Original string `json:"original"`
	Preview  string `json:"preview"`
	X96      string `json:"x96"`
	X48      string `json:"x48"`
}
