package shikimori_test

import (
	"context"
	"testing"
)

func TestVideo(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Video(context.Background(), 5114)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestVideoNullName(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Video(context.Background(), 23)

	NoError(t, err)
	NotEmpty(t, resp)
}
