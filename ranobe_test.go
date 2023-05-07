package shikimori_test

import (
	"context"
	"testing"
)

func TestRanobe(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Ranobe(context.Background(), 14893, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestRanobeRoles(t *testing.T) {
	t.Parallel()

	resp, err := shiki.RanobeRoles(context.Background(), 14893, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestRanobeSimilar(t *testing.T) {
	t.Parallel()

	resp, err := shiki.RanobeSimilar(context.Background(), 14893, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestRanobeRelated(t *testing.T) {
	t.Parallel()

	resp, err := shiki.RanobeRelated(context.Background(), 14893, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestRanobeFranchise(t *testing.T) {
	t.Parallel()

	resp, err := shiki.RanobeFranchise(context.Background(), 14893, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestRanobeExternalLinks(t *testing.T) {
	t.Parallel()

	resp, err := shiki.RanobeExternalLinks(context.Background(), 14893, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestRanobeTopics(t *testing.T) {
	t.Parallel()

	resp, err := shiki.RanobeTopics(context.Background(), 14893, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}
