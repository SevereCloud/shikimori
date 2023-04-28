package shikimori_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStudios(t *testing.T) {
	t.Parallel()

	resp, err := shiki.Studios(context.Background(), nil)

	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}
