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

type ForumsParams struct{}

func (s *API) Forums(ctx context.Context, params *ForumsParams) (resp []Forum, err error) {
	err = s.get(ctx, &resp, "forums", params)

	return
}
