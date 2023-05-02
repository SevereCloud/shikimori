package shikimori_test

import (
	"context"
	"testing"

	"github.com/SevereCloud/shikimori"
)

func TestAnimes(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Animes(context.Background(), &shikimori.AnimesParams{
		Page:  1,
		Limit: 10,
		Order: "ranked",
	})

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestAnime(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Anime(context.Background(), 5114, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestAnimeRoles(t *testing.T) {
	t.Parallel()

	resp, err := shiki.AnimeRoles(context.Background(), 5114, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestAnimeSimilar(t *testing.T) {
	t.Parallel()

	resp, err := shiki.AnimeSimilar(context.Background(), 5114, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestAnimeRelated(t *testing.T) {
	t.Parallel()

	resp, err := shiki.AnimeRelated(context.Background(), 5114, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestAnimeScreenshots(t *testing.T) {
	t.Parallel()

	resp, err := shiki.AnimeScreenshots(context.Background(), 5114, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestAnimeExternalLinks(t *testing.T) {
	t.Parallel()

	resp, err := shiki.AnimeExternalLinks(context.Background(), 5114, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestAnimeTopics(t *testing.T) {
	t.Parallel()

	resp, err := shiki.AnimeTopics(context.Background(), 5114, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}
