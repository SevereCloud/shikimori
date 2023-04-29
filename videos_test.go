package shikimori_test

import (
	"context"
	"testing"
)

func TestVideo(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Video(context.Background(), 5114, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}
