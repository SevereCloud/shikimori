package shikimori_test

import (
	"context"
	"testing"
)

func TestAnime(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Anime(context.Background(), 5114, nil)

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
