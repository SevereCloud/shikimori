package shikimori_test

import (
	"context"
	"testing"
)

func TestTopicsUpdate(t *testing.T) {
	t.Parallel()

	resp, err := shiki.TopicsUpdates(context.Background(), nil)

	NoError(t, err)
	NotEmpty(t, resp)
}
