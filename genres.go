package shikimori

type Genre struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Russian string `json:"russian"`
	Kind    string `json:"kind"`
}
