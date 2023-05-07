package shikimori

import (
	"context"
	"strconv"
)

type Manga struct {
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

type MangaFull struct {
	Manga

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

type MangaParams struct{}

func (s *API) Manga(ctx context.Context, id int, params *MangaParams) (resp MangaFull, err error) {
	err = s.get(ctx, &resp, "mangas/"+strconv.Itoa(id), params)

	return
}

type MangaRole struct {
	Roles        []string       `json:"roles"`
	RolesRussian []string       `json:"roles_russian"`
	Character    *CharacterBase `json:"character"`
	Person       *PersonBase    `json:"person"`
}

type MangaRolesParams struct{}

func (s *API) MangaRoles(
	ctx context.Context, id int, params *MangaRolesParams,
) (resp []MangaRole, err error) {
	err = s.get(ctx, &resp, "mangas/"+strconv.Itoa(id)+"/roles", params)

	return
}

type MangaSimilarParams struct{}

func (s *API) MangaSimilar(ctx context.Context, id int, params *MangaSimilarParams) (resp []Manga, err error) {
	err = s.get(ctx, &resp, "mangas/"+strconv.Itoa(id)+"/similar", params)

	return
}

type MangaRelated struct {
	Relation        string      `json:"relation"`
	RelationRussian string      `json:"relation_russian"`
	Anime           Anime       `json:"anime"`
	Manga           interface{} `json:"manga"`
}

type MangaRelatedParams struct{}

func (s *API) MangaRelated(ctx context.Context, id int, params *MangaRelatedParams) (resp []MangaRelated, err error) {
	err = s.get(ctx, &resp, "mangas/"+strconv.Itoa(id)+"/related", params)

	return
}

type MangaFranchiseParams struct{}

func (s *API) MangaFranchise(ctx context.Context, id int, params *MangaFranchiseParams) (
	resp Franchise,
	err error,
) {
	err = s.get(ctx, &resp, "mangas/"+strconv.Itoa(id)+"/franchise", params)

	return
}

type MangaExternalLinksParams struct{}

func (s *API) MangaExternalLinks(
	ctx context.Context, id int, params *MangaExternalLinksParams,
) (resp []ExternalLink, err error) {
	err = s.get(ctx, &resp, "mangas/"+strconv.Itoa(id)+"/external_links", params)

	return
}

type MangaTopicsParams struct{}

func (s *API) MangaTopics(
	ctx context.Context, id int, params *MangaTopicsParams,
) (resp []Topic, err error) {
	err = s.get(ctx, &resp, "mangas/"+strconv.Itoa(id)+"/topics", params)

	return
}
