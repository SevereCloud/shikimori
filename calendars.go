package shikimori

import (
	"context"
	"time"
)

type CalendarItem struct {
	NextEpisode   int       `json:"next_episode"`
	NextEpisodeAt time.Time `json:"next_episode_at"`
	Duration      int       `json:"duration"`
	Anime         Anime     `json:"anime"`
}

type CalendarParams struct {
	// Set to true to discard hentai, yaoi and yuri
	Censored bool `json:"censored,omitempty"`
}

func (s *API) Calendar(ctx context.Context, params *CalendarParams) (resp []CalendarItem, err error) {
	err = s.get(ctx, &resp, "calendar", params)

	return
}
