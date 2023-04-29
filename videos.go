package shikimori

import (
	"context"
	"strconv"
)

type Video struct {
	ID        int     `json:"id"`
	URL       string  `json:"url"`
	ImageURL  string  `json:"image_url"`
	PlayerURL string  `json:"player_url"`
	Name      *string `json:"name,omitempty"`
	Kind      string  `json:"kind"`
	Hosting   string  `json:"hosting"`
}

func (s *API) Video(ctx context.Context, id int) (resp []Video, err error) {
	err = s.get(ctx, &resp, "animes/"+strconv.Itoa(id)+"/videos", nil)

	return
}
