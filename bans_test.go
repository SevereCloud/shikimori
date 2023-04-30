package shikimori_test

import (
	"context"
	"testing"

	"github.com/SevereCloud/shikimori"
)

func TestBans(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Bans(context.Background(), nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestBansWithParams(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Bans(context.Background(), &shikimori.BansParams{
		Page:  1,
		Limit: 30,
	})

	NoError(t, err)
	NotEmpty(t, resp)
}
