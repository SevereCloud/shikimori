package shikimori

import (
	"context"
	"time"
)

type Comment struct {
	ID              int       `json:"id"`
	CommentableID   int       `json:"commentable_id"`
	CommentableType string    `json:"commentable_type"`
	Body            string    `json:"body"`
	UserID          int       `json:"user_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	IsOfftopic      bool      `json:"is_offtopic"`
}

type Ban struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	Comment         Comment   `json:"comment"`
	ModeratorID     int       `json:"moderator_id"`
	Reason          string    `json:"reason"`
	CreatedAt       time.Time `json:"created_at"`
	DurationMinutes int       `json:"duration_minutes"`
	User            User      `json:"user"`
	Moderator       User      `json:"moderator"`
}

type BansParams struct{}

func (s *API) Bans(ctx context.Context, params *BansParams) (resp []Ban, err error) {
	err = s.get(ctx, &resp, "bans", params)

	return
}
