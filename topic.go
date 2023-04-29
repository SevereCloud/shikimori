package shikimori

import (
	"context"
	"time"
)

type TopicsUpdate struct {
	ID        int       `json:"id"`
	Linked    Anime     `json:"linked"`
	Event     string    `json:"event"`
	Episode   int       `json:"episode"`
	CreatedAt time.Time `json:"created_at"`
	URL       string    `json:"url"`
}

type TopicsUpdatesParams struct {
	// Must be a number between 1 and 100000.
	Page int `json:"page,omitempty"`

	// 30 maximum.
	Limit int `json:"limit,omitempty"`
}

func (s *API) TopicsUpdates(ctx context.Context, params *BansParams) (resp []TopicsUpdate, err error) {
	err = s.get(ctx, &resp, "topics/updates", params)

	return
}
