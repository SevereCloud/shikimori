package shikimori_test

import (
	"context"
	"testing"
)

func TestBans(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Bans(context.Background(), nil)

	NoError(t, err)
	NotEmpty(t, resp)
}
