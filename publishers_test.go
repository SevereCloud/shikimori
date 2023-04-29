package shikimori_test

import (
	"context"
	"testing"
)

func TestPublishers(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Publishers(context.Background())

	NoError(t, err)
	NotEmpty(t, resp)
}
