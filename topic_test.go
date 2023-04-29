package shikimori_test

import (
	"context"
	"testing"

	"github.com/SevereCloud/shikimori"
)

func TestTopicsUpdates(t *testing.T) {
	t.Parallel()

	resp, err := shiki.TopicsUpdates(context.Background(), nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestTopicsHot(t *testing.T) {
	t.Parallel()

	resp, err := shiki.TopicsHot(context.Background(), &shikimori.TopicsHotParams{
		Limit: 10,
	})

	NoError(t, err)
	NotEmpty(t, resp)
}
