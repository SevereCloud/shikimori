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
