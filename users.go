package shikimori

import (
	"context"
	"strconv"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Nickname     string    `json:"nickname"`
	Avatar       string    `json:"avatar"`
	Image        UserImage `json:"image"`
	LastOnlineAt time.Time `json:"last_online_at"`
	URL          string    `json:"url"`
}

type UserImage struct {
	X160 string `json:"x160"`
	X148 string `json:"x148"`
	X80  string `json:"x80"`
	X64  string `json:"x64"`
	X48  string `json:"x48"`
	X32  string `json:"x32"`
	X16  string `json:"x16"`
}

type UsersAnimeRate struct {
	ID        int       `json:"id"`
	Score     int       `json:"score"`
	Status    string    `json:"status"`
	Text      *string   `json:"text,omitempty"`
	Episodes  *int      `json:"episodes,omitempty"`
	Chapters  *int      `json:"chapters,omitempty"`
	Volumes   *int      `json:"volumes,omitempty"`
	TextHTML  string    `json:"text_html"`
	Rewatches int       `json:"rewatches"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
	Anime     *Anime    `json:"anime,omitempty"`
	// Can represent Manga and Ranobe structs
	Manga *interface{} `json:"manga,omitempty"`
}

type UsersAnimeRateParams struct {
	// необязательно
	// Validations:
	//
	// Must be a number between 1 and 100000.
	Page int `json:"page,omitempty"`

	// необязательно
	// 5000 maximum
	//
	// Validations:
	Limit int `json:"limit,omitempty"`

	Status UserRateStatus `json:"status,omitempty"`

	// Set to true to discard hentai, yaoi and yuri
	Censored bool `json:"censored,omitempty"`
}

func (s *API) UsersAnimeRates(
	ctx context.Context, id int, params *UsersAnimeRateParams,
) (resp []UsersAnimeRate, err error) {
	err = s.get(ctx, &resp, "users/"+strconv.Itoa(id)+"/anime_rates", params)

	return
}
