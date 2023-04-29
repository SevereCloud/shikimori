package shikimori

import (
	"context"
)

type Publisher struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *API) Publishers(ctx context.Context) (resp []Publisher, err error) {
	err = s.get(ctx, &resp, "publishers", nil)

	return
}
