package shikimori

import (
	"context"
	"strconv"
	"time"
)

type CharacterBase struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Russian string `json:"russian"`
	Image   Image  `json:"image"`
	URL     string `json:"url"`
}

type Character struct {
	CharacterBase
	Altname           string           `json:"altname"`
	Japanese          string           `json:"japanese"`
	Description       *string          `json:"description"`
	DescriptionHTML   string           `json:"description_html"`
	DescriptionSource *string          `json:"description_source"`
	Favoured          bool             `json:"favoured"`
	ThreadID          *int             `json:"thread_id"`
	TopicID           *int             `json:"topic_id"`
	UpdatedAt         time.Time        `json:"updated_at"`
	Seyu              []PersonBase     `json:"seyu"`
	Animes            []CharacterAnime `json:"animes"`
	Mangas            []interface{}    `json:"mangas"`
}

type CharacterParams struct{}

func (s *API) Character(ctx context.Context, id int, params *CharacterParams) (resp Character, err error) {
	err = s.get(ctx, &resp, "characters/"+strconv.Itoa(id), params)

	return
}
