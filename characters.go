package shikimori

type CharacterBase struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Russian string `json:"russian"`
	Image   Image  `json:"image"`
	URL     string `json:"url"`
}
