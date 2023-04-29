package shikimori

import "context"

type Genre struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Russian string `json:"russian"`
	Kind    string `json:"kind"`
}

func (s *API) Genres(ctx context.Context) (resp []Genre, err error) {
	err = s.get(ctx, &resp, "genres", nil)

	return
}
