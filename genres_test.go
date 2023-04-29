package shikimori_test

import (
	"context"
	"testing"
)

func TestGenres(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Genres(context.Background())

	NoError(t, err)
	NotEmpty(t, resp)
}
