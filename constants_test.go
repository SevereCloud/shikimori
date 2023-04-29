package shikimori_test

import (
	"context"
	"testing"
)

func TestConstantsAnime(t *testing.T) {
	t.Parallel()

	resp, err := shiki.ConstantsAnime(context.Background(), nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestConstantsManga(t *testing.T) {
	t.Parallel()

	resp, err := shiki.ConstantsManga(context.Background(), nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestConstantsUserRate(t *testing.T) {
	t.Parallel()

	resp, err := shiki.ConstantsUserRate(context.Background(), nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestConstantsClub(t *testing.T) {
	t.Parallel()

	resp, err := shiki.ConstantsClub(context.Background(), nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestConstantsSmileys(t *testing.T) {
	t.Parallel()

	resp, err := shiki.ConstantsSmileys(context.Background(), nil)

	NoError(t, err)
	NotEmpty(t, resp)
}
