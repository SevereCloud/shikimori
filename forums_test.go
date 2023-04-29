package shikimori_test

import (
	"context"
	"testing"
)

func TestForums(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Forums(context.Background())

	NoError(t, err)
	NotEmpty(t, resp)
}
