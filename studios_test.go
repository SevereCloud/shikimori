package shikimori_test

import (
	"context"
	"testing"
)

func TestStudios(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Studios(context.Background(), nil)

	NoError(t, err)
	NotEmpty(t, resp)
}
