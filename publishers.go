package shikimori

import (
	"context"
)

type Publisher struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PublishersParams struct{}

func (s *API) Publishers(ctx context.Context, params *PublishersParams) (resp []Publisher, err error) {
	err = s.get(ctx, &resp, "publishers", params)

	return
}
