package shikimori_test

import (
	"context"
	"testing"
)

func TestCharacter(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Character(context.Background(), 30, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}
