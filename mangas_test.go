package shikimori_test

import (
	"context"
	"testing"
)

func TestManga(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Manga(context.Background(), 27, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestMangaRoles(t *testing.T) {
	t.Parallel()

	resp, err := shiki.MangaRoles(context.Background(), 27, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestMangaSimilar(t *testing.T) {
	t.Parallel()

	resp, err := shiki.MangaSimilar(context.Background(), 27, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestMangaRelated(t *testing.T) {
	t.Parallel()

	resp, err := shiki.MangaRelated(context.Background(), 27, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestMangaExternalLinks(t *testing.T) {
	t.Parallel()

	resp, err := shiki.MangaExternalLinks(context.Background(), 27, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestMangaTopics(t *testing.T) {
	t.Parallel()

	resp, err := shiki.MangaTopics(context.Background(), 27, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}
