package shikimori

import (
	"context"
)

type Forum struct {
	ID        int    `json:"id"`
	Position  int    `json:"position"`
	Name      string `json:"name"`
	Permalink string `json:"permalink"`
	URL       string `json:"url"`
}

func (s *API) Forums(ctx context.Context) (resp []Forum, err error) {
	err = s.get(ctx, &resp, "forums", nil)

	return
}
