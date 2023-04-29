package shikimori

import (
	"context"
)

type Studio struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	FilteredName string  `json:"filtered_name"`
	Real         bool    `json:"real"`
	Image        *string `json:"image,omitempty"`
}

func (s *API) Studios(ctx context.Context) (resp []Studio, err error) {
	err = s.get(ctx, &resp, "studios", nil)

	return
}
