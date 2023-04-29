package shikimori

type Video struct {
	ID        int     `json:"id"`
	URL       string  `json:"url"`
	ImageURL  string  `json:"image_url"`
	PlayerURL string  `json:"player_url"`
	Name      *string `json:"name"`
	Kind      string  `json:"kind"`
	Hosting   string  `json:"hosting"`
}
