package shikimori

import (
	"context"
	"strconv"
)

type Ranobe struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Russian    string `json:"russian"`
	Image      Image  `json:"image"`
	URL        string `json:"url"`
	Kind       string `json:"kind"`
	Score      string `json:"score"`
	Status     string `json:"status"`
	Volumes    int    `json:"volumes"`
	Chapters   int    `json:"chapters"`
	AiredOn    string `json:"aired_on"`
	ReleasedOn string `json:"released_on"`
}

type RanobeFull struct {
	Ranobe

	English            []string            `json:"english"`
	Japanese           []string            `json:"japanese"`
	Synonyms           []string            `json:"synonyms"`
	LicenseNameRu      *string             `json:"license_name_ru"`
	Description        string              `json:"description"`
	DescriptionHTML    string              `json:"description_html"`
	DescriptionSource  *string             `json:"description_source"`
	Franchise          *string             `json:"franchise"`
	Favoured           bool                `json:"favoured"`
	Anons              bool                `json:"anons"`
	Ongoing            bool                `json:"ongoing"`
	ThreadID           *int                `json:"thread_id"`
	TopicID            *int                `json:"topic_id"`
	MyanimelistID      int                 `json:"myanimelist_id"`
	RatesScoresStats   []RatesScoresStat   `json:"rates_scores_stats"`
	RatesStatusesStats []RatesStatusesStat `json:"rates_statuses_stats"`
	Licensors          []string            `json:"licensors"`
	Genres             []Genre             `json:"genres"`
	Publishers         []Publisher         `json:"publishers"`
	UserRate           *UserRate           `json:"user_rate"`
}

type RanobeParams struct{}

func (s *API) Ranobe(ctx context.Context, id int, params *RanobeParams) (resp RanobeFull, err error) {
	err = s.get(ctx, &resp, "ranobe/"+strconv.Itoa(id), params)

	return
}

type RanobeRole struct {
	Roles        []string       `json:"roles"`
	RolesRussian []string       `json:"roles_russian"`
	Character    *CharacterBase `json:"character"`
	Person       *PersonBase    `json:"person"`
}

type RanobeRolesParams struct{}

func (s *API) RanobeRoles(
	ctx context.Context, id int, params *RanobeRolesParams,
) (resp []RanobeRole, err error) {
	err = s.get(ctx, &resp, "ranobe/"+strconv.Itoa(id)+"/roles", params)

	return
}

type RanobeSimilarParams struct{}

func (s *API) RanobeSimilar(ctx context.Context, id int, params *RanobeSimilarParams) (resp []Ranobe, err error) {
	err = s.get(ctx, &resp, "ranobe/"+strconv.Itoa(id)+"/similar", params)

	return
}

type RanobeRelated struct {
	Relation        string      `json:"relation"`
	RelationRussian string      `json:"relation_russian"`
	Anime           Anime       `json:"anime"`
	Manga           interface{} `json:"manga"`
}

type RanobeRelatedParams struct{}

func (s *API) RanobeRelated(ctx context.Context, id int, params *RanobeRelatedParams) (
	resp []RanobeRelated,
	err error,
) {
	err = s.get(ctx, &resp, "ranobe/"+strconv.Itoa(id)+"/related", params)

	return
}

type RanobeFranchiseParams struct{}

func (s *API) RanobeFranchise(ctx context.Context, id int, params *RanobeFranchiseParams) (
	resp Franchise,
	err error,
) {
	err = s.get(ctx, &resp, "ranobe/"+strconv.Itoa(id)+"/franchise", params)

	return
}

type RanobeExternalLinksParams struct{}

func (s *API) RanobeExternalLinks(ctx context.Context, id int, params *RanobeExternalLinksParams) (
	resp []ExternalLink,
	err error,
) {
	err = s.get(ctx, &resp, "ranobe/"+strconv.Itoa(id)+"/external_links", params)

	return
}

type RanobeTopicsParams struct{}

func (s *API) RanobeTopics(
	ctx context.Context, id int, params *RanobeTopicsParams,
) (resp []Topic, err error) {
	err = s.get(ctx, &resp, "ranobe/"+strconv.Itoa(id)+"/topics", params)

	return
}
