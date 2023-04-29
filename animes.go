package shikimori

import (
	"context"
	"strconv"
	"time"
)

type RatesScoresStat struct {
	Name  int `json:"name"`
	Value int `json:"value"`
}

type RatesStatusesStat struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Anime struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	Russian       string     `json:"russian"`
	Image         AnimeImage `json:"image"`
	URL           string     `json:"url"`
	Kind          string     `json:"kind"`
	Score         string     `json:"score"`
	Status        string     `json:"status"`
	Episodes      int        `json:"episodes"`
	EpisodesAired int        `json:"episodes_aired"`
	AiredOn       string     `json:"aired_on"`
	ReleasedOn    *string    `json:"released_on"`
}

type AnimeFull struct {
	Anime

	Rating             string              `json:"rating"`
	English            []string            `json:"english"`
	Japanese           []string            `json:"japanese"`
	Synonyms           []string            `json:"synonyms"`
	LicenseNameRu      *string             `json:"license_name_ru"`
	Duration           int                 `json:"duration"`
	Description        string              `json:"description"`
	DescriptionHTML    string              `json:"description_html"`
	DescriptionSource  interface{}         `json:"description_source"`
	Franchise          string              `json:"franchise"`
	Favoured           bool                `json:"favoured"`
	Anons              bool                `json:"anons"`
	Ongoing            bool                `json:"ongoing"`
	ThreadID           int                 `json:"thread_id"`
	TopicID            int                 `json:"topic_id"`
	MyanimelistID      int                 `json:"myanimelist_id"`
	RatesScoresStats   []RatesScoresStat   `json:"rates_scores_stats"`
	RatesStatusesStats []RatesStatusesStat `json:"rates_statuses_stats"`
	UpdatedAt          time.Time           `json:"updated_at"`
	NextEpisodeAt      *time.Time          `json:"next_episode_at"`
	Fansubbers         []string            `json:"fansubbers"`
	Fandubbers         []string            `json:"fandubbers"`
	Licensors          []string            `json:"licensors"`
	Genres             []Genre             `json:"genres"`
	Studios            []Studio            `json:"studios"`
	Videos             []Video             `json:"videos"`
	Screenshots        []Screenshot        `json:"screenshots"`
	UserRate           *UserRate           `json:"user_rate"`
}

type AnimeImage struct {
	Original string `json:"original"`
	Preview  string `json:"preview"`
	X96      string `json:"x96"`
	X48      string `json:"x48"`
}

type AnimeParams struct{}

func (s *API) Anime(ctx context.Context, id int, params *AnimeParams) (resp AnimeFull, err error) {
	err = s.get(ctx, &resp, "animes/"+strconv.Itoa(id), params)

	return
}

type AnimeScreenshotsParams struct{}

func (s *API) AnimeScreenshots(
	ctx context.Context, id int, params *AnimeScreenshotsParams,
) (resp []Screenshot, err error) {
	err = s.get(ctx, &resp, "animes/"+strconv.Itoa(id)+"/screenshots", params)

	return
}

type ExternalLinks struct {
	ID         int        `json:"id"`
	Kind       string     `json:"kind"`
	URL        string     `json:"url"`
	Source     string     `json:"source"`
	EntryID    int        `json:"entry_id"`
	EntryType  string     `json:"entry_type"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	ImportedAt *time.Time `json:"imported_at"`
}

type AnimeExternalLinksParams struct{}

func (s *API) AnimeExternalLinks(
	ctx context.Context, id int, params *AnimeExternalLinksParams,
) (resp []Screenshot, err error) {
	err = s.get(ctx, &resp, "animes/"+strconv.Itoa(id)+"/external_links", params)

	return
}
