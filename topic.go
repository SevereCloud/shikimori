package shikimori

import (
	"context"
	"time"
)

// Can be Anime, Manga structs
type TopicLinked struct {
	Anime
	// Manga
}

type Topic struct {
	ID                int          `json:"id"`
	TopicTitle        string       `json:"topic_title"`
	Body              string       `json:"body"`
	HTMLBody          string       `json:"html_body"`
	HTMLFooter        string       `json:"html_footer"`
	CreatedAt         time.Time    `json:"created_at"`
	CommentsCount     int          `json:"comments_count"`
	Forum             Forum        `json:"forum"`
	User              User         `json:"user"`
	Type              string       `json:"type"`
	LinkedID          int          `json:"linked_id"`
	LinkedType        string       `json:"linked_type"`
	Linked            *TopicLinked `json:"linked,omitempty"`
	Viewed            bool         `json:"viewed"`
	LastCommentViewed *bool        `json:"last_comment_viewed,omitempty"`
	Event             *string      `json:"event,omitempty"`
	Episode           *int         `json:"episode,omitempty"`
}

type TopicsUpdate struct {
	ID        int         `json:"id"`
	Linked    TopicLinked `json:"linked"`
	Event     *string     `json:"event,omitempty"`
	Episode   *int        `json:"episode,omitempty"`
	CreatedAt time.Time   `json:"created_at"`
	URL       string      `json:"url"`
}

type TopicsUpdatesParams struct {
	// Must be a number between 1 and 100000.
	Page int `json:"page,omitempty"`

	// 30 maximum.
	Limit int `json:"limit,omitempty"`
}

func (s *API) TopicsUpdates(ctx context.Context, params *TopicsUpdatesParams) (resp []TopicsUpdate, err error) {
	err = s.get(ctx, &resp, "topics/updates", params)

	return
}

type TopicsHotParams struct {
	// 10 maximum.
	Limit int `json:"limit,omitempty"`
}

func (s *API) TopicsHot(ctx context.Context, params *TopicsHotParams) (resp []Topic, err error) {
	err = s.get(ctx, &resp, "topics/hot", params)

	return
}
