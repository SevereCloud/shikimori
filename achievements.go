package shikimori

import (
	"context"
	"time"
)

type Achievement struct {
	ID        int       `json:"id"`
	NekoID    string    `json:"neko_id"`
	Level     int       `json:"level"`
	Progress  int       `json:"progress"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AchievementsParams struct {
	UserID int `json:"user_id"`
}

// Achievements return list user achievements.
func (s *API) Achievements(ctx context.Context, params *AchievementsParams) (resp []Achievement, err error) {
	err = s.get(ctx, &resp, "achievements", params)

	return
}
