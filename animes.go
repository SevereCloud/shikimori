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
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Russian       string  `json:"russian"`
	Image         Image   `json:"image"`
	URL           string  `json:"url"`
	Kind          *string `json:"kind,omitempty"`
	Score         string  `json:"score"`
	Status        string  `json:"status"`
	Episodes      int     `json:"episodes"`
	EpisodesAired int     `json:"episodes_aired"`
	AiredOn       *string `json:"aired_on,omitempty"`
	ReleasedOn    *string `json:"released_on,omitempty"`
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

type AnimeOrder string

const (
	AnimeOrderID            AnimeOrder = "id"
	AnimeOrderIDDesc        AnimeOrder = "id_desc"
	AnimeOrderRanked        AnimeOrder = "ranked"
	AnimeOrderKind          AnimeOrder = "kind"
	AnimeOrderPopularity    AnimeOrder = "popularity"
	AnimeOrderName          AnimeOrder = "name"
	AnimeOrderAiredOn       AnimeOrder = "aired_on"
	AnimeOrderEpisodes      AnimeOrder = "episodes"
	AnimeOrderStatus        AnimeOrder = "status"
	AnimeOrderRandom        AnimeOrder = "random"
	AnimeOrderCreatedAt     AnimeOrder = "created_at"
	AnimeOrderCreatedAtDesc AnimeOrder = "created_at_desc"
)

type AnimeDuration string

const (
	// Less than 10 minutes.
	AnimeDurationsS AnimeDuration = "S"
	// Less than 30 minutes.
	AnimeDurationsD AnimeDuration = "D"
	// More than 30 minutes.
	AnimeDurationsF AnimeDuration = "F"
)

type AnimeRating string

const (
	// No rating.
	AnimeRatingNone AnimeRating = "none"
	// All ages.
	AnimeRatingG AnimeRating = "g"
	// Children.
	AnimeRatingPG AnimeRating = "pg"
	// Teens 13 or older.
	AnimeRatingPG13 AnimeRating = "pg_13"
	// Recommended (violence & profanity).
	AnimeRatingR AnimeRating = "r"
	// Mild Nudity (may also contain violence & profanity).
	AnimeRatingRPlus AnimeRating = "r_plus"
	// Hentai (extreme sexual content/nudity).
	AnimeRatingRX AnimeRating = "rx"
)

type AnimeMyList string

const (
	AnimeMyListPlanned    AnimeMyList = "planned"
	AnimeMyListWatching   AnimeMyList = "watching"
	AnimeMyListRewatching AnimeMyList = "rewatching"
	AnimeMyListCompleted  AnimeMyList = "completed"
	AnimeMyListOnHold     AnimeMyList = "on_hold"
	AnimeMyListDropped    AnimeMyList = "dropped"
)

type AnimesParams struct {
	// Must be a number between 1 and 100000.
	Page int `json:"page,omitempty"`

	// 50 maximum. Must be a number.
	Limit int `json:"limit,omitempty"`

	// Must be one of: id, id_desc, ranked, kind, popularity, name,
	// aired_on, episodes, status, random, ranked_random,
	// ranked_shiki, created_at, created_at_desc.
	Order AnimeOrder `json:"order,omitempty"`

	// Must be one of: tv, movie, ova, ona, special, music, tv_13, tv_24, tv_48.
	Kind AnimeKind `json:"kind,omitempty"`

	// Must be one of: anons, ongoing, released.
	Status AnimeStatus `json:"status,omitempty"`

	// Examples: summer_2017, 2016, 2014_2016, 199x.
	Season string `json:"season,omitempty"`

	// Minimal anime score. Must be a number.
	Score int `json:"score,omitempty"`

	// Must be one of: S, D, F.
	Duration AnimeDuration `json:"duration,omitempty"`

	// Must be one of: none, g, pg, pg_13, r, r_plus, rx.
	Rating AnimeRating `json:"rating,omitempty"`

	// List of genre ids separated by comma.
	Genre string `json:"genre,omitempty"`

	// List of studio ids separated by comma.
	Studio string `json:"studio,omitempty"`

	// List of franchises separated by comma.
	Franchise string `json:"franchise,omitempty"`

	// Set to false to allow hentai, yaoi and yuri. Must be one of: true, false.
	Censored bool `json:"censored,omitempty"`

	// Status of manga in current user list.
	// Must be one of: planned, watching, rewatching, completed, on_hold, dropped.
	MyList AnimeMyList `json:"mylist,omitempty"`

	// List of anime ids separated by comma.
	IDs string `json:"ids,omitempty"`

	// List of anime ids separated by comma.
	ExcludeIDs string `json:"exclude_ids,omitempty"`

	// Search phrase to filter animes by name. Must be a String.
	Search string `json:"search,omitempty"`
}

func (s *API) Animes(ctx context.Context, params *AnimesParams) (resp []Anime, err error) {
	err = s.get(ctx, &resp, "animes", params)

	return
}

type AnimeParams struct{}

func (s *API) Anime(ctx context.Context, id int, params *AnimeParams) (resp AnimeFull, err error) {
	err = s.get(ctx, &resp, "animes/"+strconv.Itoa(id), params)

	return
}

type AnimeRole struct {
	Roles        []string       `json:"roles"`
	RolesRussian []string       `json:"roles_russian"`
	Character    *CharacterBase `json:"character"`
	Person       *PersonBase    `json:"person"`
}

type AnimeRolesParams struct{}

func (s *API) AnimeRoles(
	ctx context.Context, id int, params *AnimeRolesParams,
) (resp []AnimeRole, err error) {
	err = s.get(ctx, &resp, "animes/"+strconv.Itoa(id)+"/roles", params)

	return
}

type AnimeSimilarParams struct{}

func (s *API) AnimeSimilar(ctx context.Context, id int, params *AnimeSimilarParams) (resp []Anime, err error) {
	err = s.get(ctx, &resp, "animes/"+strconv.Itoa(id)+"/similar", params)

	return
}

type AnimeRelated struct {
	Relation        string      `json:"relation"`
	RelationRussian string      `json:"relation_russian"`
	Anime           Anime       `json:"anime"`
	Manga           interface{} `json:"manga"`
}

type AnimeRelatedParams struct{}

func (s *API) AnimeRelated(ctx context.Context, id int, params *AnimeRelatedParams) (resp []AnimeRelated, err error) {
	err = s.get(ctx, &resp, "animes/"+strconv.Itoa(id)+"/related", params)

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

type AnimeExternalLinksParams struct{}

func (s *API) AnimeExternalLinks(
	ctx context.Context, id int, params *AnimeExternalLinksParams,
) (resp []ExternalLinks, err error) {
	err = s.get(ctx, &resp, "animes/"+strconv.Itoa(id)+"/external_links", params)

	return
}
