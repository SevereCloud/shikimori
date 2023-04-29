package shikimori

import (
	"context"
	"strconv"
	"time"
)

type UserRate struct {
	ID         int            `json:"id"`
	UserID     int            `json:"user_id"`
	TargetID   int            `json:"target_id"`
	TargetType string         `json:"target_type"`
	Score      int            `json:"score"`
	Status     UserRateStatus `json:"status"`
	Rewatches  int            `json:"rewatches"`
	Episodes   int            `json:"episodes"`
	Volumes    int            `json:"volumes"`
	Chapters   int            `json:"chapters"`
	Text       *string        `json:"text"`
	TextHTML   string         `json:"text_html"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

type UserRateParams struct{}

func (s *API) UserRate(ctx context.Context, id int, params *UserRateParams) (resp []UserRate, err error) {
	err = s.get(ctx, &resp, "v2/user_rates/"+strconv.Itoa(id), params)

	return
}
