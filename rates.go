package shikimori

import (
	"context"
	"strconv"
	"time"
)

type UserRateTargetType string

const (
	UserRateTargetAnime UserRateTargetType = "Anime"
	UserRateTargetManga UserRateTargetType = "Manga"
)

type UserRate struct {
	ID         int                 `json:"id"`
	UserID     *int                `json:"user_id,omitempty"`
	TargetID   *int                `json:"target_id,omitempty"`
	TargetType *UserRateTargetType `json:"target_type,omitempty"`
	Score      int                 `json:"score"`
	Status     UserRateStatus      `json:"status"`
	Rewatches  int                 `json:"rewatches"`
	Episodes   *int                `json:"episodes,omitempty"`
	Volumes    *int                `json:"volumes,omitempty"`
	Chapters   *int                `json:"chapters,omitempty"`
	Text       *string             `json:"text,omitempty"`
	TextHTML   string              `json:"text_html"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}

func (s *API) UserRate(ctx context.Context, id int) (resp UserRate, err error) {
	err = s.get(ctx, &resp, "v2/user_rates/"+strconv.Itoa(id), nil)

	return
}

type UserRatesParams struct {
	UserID int `json:"user_id,omitempty"`

	TargetID int `json:"target_id,omitempty"`

	// Must be one of: Anime, Manga.
	TargetType UserRateTargetType `json:"target_type,omitempty"`

	// Must be one of: planned, watching, rewatching, completed, on_hold, dropped
	Status UserRateStatus `json:"status,omitempty"`

	// Must be a number between 1 and 100000. This field is ignored when user_id is set
	Page int `json:"page,omitempty"`

	// 1000 maximum. This field is ignored when user_id is set
	Limit int `json:"limit,omitempty"`
}

func (s *API) UserRates(ctx context.Context, params *UserRatesParams) (resp []UserRate, err error) {
	err = s.get(ctx, &resp, "v2/user_rates", params)

	return
}
