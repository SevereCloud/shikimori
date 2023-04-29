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
	Kind          *string    `json:"kind,omitempty"`
	Score         string     `json:"score"`
	Status        string     `json:"status"`
	Episodes      int        `json:"episodes"`
	EpisodesAired int        `json:"episodes_aired"`
	AiredOn       *string    `json:"aired_on,omitempty"`
	ReleasedOn    *string    `json:"released_on,omitempty"`
}

type AnimeFull struct {
	Anime

	Rating             string              `json:"rating"`
	English            []string            `json:"english"`
	Japanese           []string            `json:"japanese"`
	Synonyms           []string            `json:"synonyms"`
	LicenseNameRu      *string             `json:"license_name_ru,omitempty"`
	Duration           int                 `json:"duration"`
	Description        *string             `json:"description,omitempty"`
	DescriptionHTML    string              `json:"description_html"`
	DescriptionSource  *string             `json:"description_source,omitempty"`
	Franchise          *string             `json:"franchise,omitempty"`
	Favoured           bool                `json:"favoured"`
	Anons              bool                `json:"anons"`
	Ongoing            bool                `json:"ongoing"`
	ThreadID           *int                `json:"thread_id,omitempty"`
	TopicID            *int                `json:"topic_id,omitempty"`
	MyanimelistID      int                 `json:"myanimelist_id"`
	RatesScoresStats   []RatesScoresStat   `json:"rates_scores_stats"`
	RatesStatusesStats []RatesStatusesStat `json:"rates_statuses_stats"`
	UpdatedAt          time.Time           `json:"updated_at"`
	NextEpisodeAt      *time.Time          `json:"next_episode_at,omitempty"`
	Fansubbers         []string            `json:"fansubbers"`
	Fandubbers         []string            `json:"fandubbers"`
	Licensors          []string            `json:"licensors"`
	Genres             []Genre             `json:"genres"`
	Studios            []Studio            `json:"studios"`
	Videos             []Video             `json:"videos"`
	Screenshots        []Screenshot        `json:"screenshots"`
	UserRate           *UserRate           `json:"user_rate,omitempty"`
}

type AnimeImage struct {
	Original string `json:"original"`
	Preview  string `json:"preview"`
	X96      string `json:"x96"`
	X48      string `json:"x48"`
}

type AnimeParams struct {
	Page       int    `json:"page,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Order      string `json:"order,omitempty"`
	Kind       string `json:"kind,omitempty"`
	Status     string `json:"status,omitempty"`
	Season     string `json:"season,omitempty"`
	Score      int    `json:"score,omitempty"`
	Duration   string `json:"duration,omitempty"`
	Rating     string `json:"rating,omitempty"`
	Genre      string `json:"genre,omitempty"`
	Studio     string `json:"studio,omitempty"`
	Franchise  string `json:"franchise,omitempty"`
	Censored   bool   `json:"censored,omitempty"`
	MyList     string `json:"mylist,omitempty"`
	IDs        string `json:"ids,omitempty"`
	ExcludeIDs string `json:"exclude_ids,omitempty"`
	Search     string `json:"search,omitempty"`
}

func (s *API) Animes(ctx context.Context, params *AnimeParams) (resp []Anime, err error) {
	err = s.get(ctx, &resp, "animes", params)

	return
}

func (s *API) Anime(ctx context.Context, id int) (resp AnimeFull, err error) {
	err = s.get(ctx, &resp, "animes/"+strconv.Itoa(id), nil)

	return
}

func (s *API) AnimeScreenshots(ctx context.Context, id int) (resp []Screenshot, err error) {
	err = s.get(ctx, &resp, "animes/"+strconv.Itoa(id)+"/screenshots", nil)

	return
}

type ExternalLinks struct {
	ID         *int       `json:"id,omitempty"`
	Kind       string     `json:"kind"`
	URL        string     `json:"url"`
	Source     string     `json:"source"`
	EntryID    int        `json:"entry_id"`
	EntryType  string     `json:"entry_type"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	ImportedAt *time.Time `json:"imported_at,omitempty"`
}

func (s *API) AnimeExternalLinks(ctx context.Context, id int) (resp []Screenshot, err error) {
	err = s.get(ctx, &resp, "animes/"+strconv.Itoa(id)+"/external_links", nil)

	return
}
